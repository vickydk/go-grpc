package protobuf

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/golang/glog"
	hSvc "github.com/vickydk/go-grpc/pkg/protobuf/helloworld"
	h "github.com/vickydk/go-grpc/pkg/utl/grpc"
	"github.com/vickydk/go-grpc/pkg/utl/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
)

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "self", "go-rpc", path)
}

func Start() error {
	var (
		caCert          = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		tlsCert         = flag.String("tls-cert", withConfigDir("oms.pem"), "TLS server certificate.")
		tlsKey          = flag.String("tls-key", withConfigDir("oms-key.pem"), "TLS server key.")
	)
	flag.Parse()

	lis, err := net.Listen("tcp", "127.0.0.1:50058")
	if err != nil {
		fmt.Println("failed to listen: %v", err)
		return err
	}

	defer func() {
		if err := lis.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", "tcp", "127.0.0.1:50058", err)
		}
	}()

	cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	if err != nil {
		fmt.Println("err: ",err)
	}
	rawCaCert, err := ioutil.ReadFile(*caCert)
	if err != nil {
		fmt.Println("err: ",err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(rawCaCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	h.RegisterGreeterServer(grpcServer, hSvc.NewGrpcServer())

	healthServer := health.NewServer()
	healthServer.SetServingStatus("grpc.health.v1.helloservice", 1)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	reflection.Register(grpcServer)
	server.StartGrpc(grpcServer, lis)

	return nil
}
