package server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func StartGrpc(grpcServer *grpc.Server, lis net.Listener) {
	// Start server
	fmt.Println("Serving gRPC on tcp://127.0.0.1:50058")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println("StartGrpc - failed to serve: ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Stopping gRPC Server")
}

