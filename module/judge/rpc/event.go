package rpc

import (
	"context"
	"github.com/GeorgeWang1994/cicada/module/judge/judge"
	"github.com/GeorgeWang1994/cicada/module/pkg/model"
	pb "github.com/GeorgeWang1994/cicada/module/proto/api/judge"
	log "github.com/sirupsen/logrus"
	"time"
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
	for _, e := range request.Events {
		if err := judge.Judge(&model.HoneypotEvent{
			ID:        e.Id,
			Proto:     e.Proto,
			Honeypot:  e.Honeypot,
			Agent:     e.Agent,
			StartTime: time.Unix(e.StartTime.Seconds, 0),
			EndTime:   time.Unix(e.EndTime.Seconds, 0),
			SrcIp:     e.SrcIp,
			SrcPort:   int(e.SrcPort),
			DestIp:    e.DestIp,
			DestPort:  int(e.DestPort),
			RiskLevel: int(e.RiskLevel),
		}); err != nil {
			log.Errorf("judge event %s error %v", e.Id, err)
			continue
		}
	}
	return &pb.Response{}, nil
}
