package rpc

import (
	"cicada/judge/cc"
	"cicada/pkg/utils/rpc"
	pb "cicada/proto/api/judge"
	"context"
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
