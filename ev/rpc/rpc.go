package rpc

import (
	"cicada/ev/cc"
	msgpackrpc "github.com/hashicorp/net-rpc-msgpackrpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func Start() {
	addr := cc.Config().Rpc.RpcAddr

	server := rpc.NewServer()
	err := server.Register(new(Event))
	if err != nil {
		log.Fatalln(err)
	}

	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalln("listen error:", e)
	} else {
		log.Println("listening", addr)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("listener accept fail:", err)
			time.Sleep(time.Duration(100) * time.Millisecond)
			continue
		}
		msgpackrpc.NewServerCodec(conn)
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
