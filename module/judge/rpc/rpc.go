package rpc

import (
	"context"
	"github.com/GeorgeWang1994/cicada/module/judge/cc"
	"github.com/GeorgeWang1994/cicada/module/pkg/utils/rpc"
	pb "github.com/GeorgeWang1994/cicada/module/proto/api/judge"
	log "github.com/sirupsen/logrus"
	"net"
)

func Start(ctx context.Context) {
	addr := cc.Config().Rpc.RpcAddr

	gs := rpc.NewGRpcServer(ctx)
	pb.RegisterJudgeServiceServer(gs, New())

	ln, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalln("listen error:", e)
	} else {
		log.Println("listening", addr)
	}

	go gs.Serve(ln)
}
