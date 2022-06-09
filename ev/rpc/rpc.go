package rpc

import (
	"context"
	"github.com/GeorgeWang1994/cicada/ev/cc"
	"github.com/GeorgeWang1994/cicada/pkg/utils/rpc"
	pb "github.com/GeorgeWang1994/cicada/proto/api/ev"
	log "github.com/sirupsen/logrus"
	"net"
)

func Start(ctx context.Context) {
	addr := cc.Config().Rpc.RpcAddr

	gs := rpc.NewGRpcServer(ctx)
	pb.RegisterEventServiceServer(gs, New())

	ln, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalln("listen error:", e)
	} else {
		log.Println("listening", addr)
	}

	go gs.Serve(ln)
}
