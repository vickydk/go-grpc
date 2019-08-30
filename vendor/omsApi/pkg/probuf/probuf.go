package probuf

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"net"
	"omsApi/pkg/domain/referential"
	rfSvc "omsApi/pkg/probuf/referential/transport"
	"omsApi/pkg/utl/config"
	dbhandler "omsApi/pkg/utl/dbhandler/mysql"
	rf "omsApi/pkg/utl/grpc/referential"
	"omsApi/pkg/utl/server"
	"os"
	"path/filepath"
)

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "omsApi", path)
}

func Start() error {
	var (
		caCert          = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		tlsCert         = flag.String("tls-cert", withConfigDir("oms.pem"), "TLS server certificate.")
		tlsKey          = flag.String("tls-key", withConfigDir("oms-key.pem"), "TLS server key.")
	)
	flag.Parse()

	db, err := dbhandler.New(config.Env.Driver, config.Env.PSN, config.Env.MaxPool, config.Env.MaxIdle)
	if err != nil {
		return err
	}

	lis, err := net.Listen(config.Env.Protocol, config.Env.PortServer)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
		return err
	}

	defer func() {
		if err := lis.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", config.Env.Protocol, config.Env.PortServer, err)
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

	rf.RegisterReferentialServiceServer(grpcServer, rfSvc.NewGrpcServer(referential.Initialize(db)))

	healthServer := health.NewServer()
	healthServer.SetServingStatus("grpc.health.v1.helloservice", 1)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	reflection.Register(grpcServer)
	server.StartGrpc(grpcServer, lis)

	return nil
}
