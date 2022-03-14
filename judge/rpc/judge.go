package rpc

import (
	pb "cicada/proto/api/judge"
	"context"
)

type Judge struct {
	pb.UnimplementedJudgeServiceServer
}

func New() *Judge {
	return &Judge{}
}

func (t *Judge) Ping(ctx context.Context, request *pb.Empty) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (t *Judge) ReceiveEvent(ctx context.Context, request *pb.ReceiveEventRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
