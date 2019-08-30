package transport

import (
	"context"
	"fmt"
	"omsApi/pkg/domain/referential"
	rf "omsApi/pkg/utl/grpc/referential"
	omsApi "omsApi/pkg/utl/model"
	"omsApi/pkg/utl/structs"
)

type grpcServer struct {
	svc referential.Service
}

func NewGrpcServer(svc referential.Service) rf.ReferentialServiceServer {
	return &grpcServer{svc: svc}
}

func (g *grpcServer) List(ctx context.Context, req *rf.ReferentialReq) (*rf.ReferentialListResp, error) {
	var res rf.ReferentialListResp
	orderReq := new(omsApi.ReferentialReq)
	structs.Merge(orderReq, req)
	result, err := g.svc.List(orderReq)
	if err != nil {
		return nil, err
	}

	var resList []*rf.ReferentialResp
	for _, eachrefs := range result {
		res := new(rf.ReferentialResp)
		structs.Merge(res, &eachrefs)
		resList = append(resList, res)
	}

	res.ReferentialList = resList

	return &res, nil
}

func (g *grpcServer) ListStream(req *rf.ReferentialReq, stream rf.ReferentialService_ListStreamServer) error {
	orderReq := new(omsApi.ReferentialReq)
	structs.Merge(orderReq, req)
	result, err := g.svc.List(orderReq)
	if err != nil {
		return err
	}

	for _, eachrefs := range result {
		res := new(rf.ReferentialResp)
		structs.Merge(res, &eachrefs)
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	fmt.Println("DONE!!")
	return nil
}
