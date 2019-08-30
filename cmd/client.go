package main

import (
	"context"
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	h "self/go-grpc/pkg/utl/grpc"
)

const (
	address     = "localhost:50052"
)

func main() {
	var (
		caCert          = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		tlsCert         = flag.String("tls-cert", withConfigDir("oms.pem"), "TLS server certificate.")
		tlsKey          = flag.String("tls-key", withConfigDir("oms-key.pem"), "TLS server key.")
	)
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	if err != nil {
		fmt.Println("err: ",err)
	}

	rawCACert, err := ioutil.ReadFile(*caCert)
	if err != nil {
		fmt.Println("err: ",err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(rawCACert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	})

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	c := h.NewGreeterClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	for i:=1; i<10; i ++ {
		goSend(c, ctx, int32(i))
	}
}

func goSend(c h.GreeterClient, ctx context.Context, id int32) {
	r, err := c.SayHello(ctx, &h.HelloRequest{Id: id})
	if err != nil {
		fmt.Println("could not greet: %v", err)
	}
	fmt.Println("Greeting: ID", id, " | RES: ", r)
}

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "self", "go-grpc", path)
}
