package rpc

import (
	"context"
	"github.com/GeorgeWang1994/cicada/module/pkg/utils/rpc"
	"github.com/GeorgeWang1994/cicada/module/portal/cc"
	log "github.com/sirupsen/logrus"
	"net"
)

func Start(ctx context.Context) {
	addr := cc.Config().Rpc.RpcAddr

	gs := rpc.NewGRpcServer(ctx)

	ln, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalln("listen error:", e)
	} else {
		log.Println("listening", addr)
	}

	go gs.Serve(ln)
}
