package rpc

import (
	"cicada/ev/cc"
	"cicada/pkg/utils/rpc"
	pb "cicada/proto/api/ev"
	"context"
	"log"
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
