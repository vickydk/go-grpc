package helloworld

import (
	"context"
	h "github.com/vickydk/go-grpc/pkg/utl/grpc"
)

type grpcServer struct {
}

func NewGrpcServer() h.GreeterServer {
	return &grpcServer{}
}

func (grpcServer) SayHello(ctx context.Context, req *h.HelloRequest) (*h.HelloReply, error) {
	var res h.HelloReply
	res.Id = req.Id
	
	return &res, nil
}