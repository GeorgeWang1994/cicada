package rpc

import (
	"cicada/judge/judge"
	"cicada/pkg/model"
	pb "cicada/proto/api/judge"
	"context"
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
			ID:         e.Id,
			Proto:      e.Proto,
			Honeypot:   e.Honeypot,
			Agent:      e.Agent,
			StartTime:  time.Unix(e.StartTime.Seconds, 0),
			EndTime:    time.Unix(e.EndTime.Seconds, 0),
			SrcIp:      e.SrcIp,
			SrcPort:    int(e.SrcPort),
			SrcMac:     e.SrcMac,
			DestIp:     e.DestIp,
			DestPort:   int(e.DestPort),
			EventTypes: e.EventTypes,
			RiskLevel:  int(e.RiskLevel),
		}); err != nil {
			log.Errorf("judge event error %s", e.Id)
		}
	}
	return &pb.Response{}, nil
}
