package gg

import (
	"github.com/GeorgeWang1994/cicada/module/agentd/cc"
	"github.com/GeorgeWang1994/cicada/module/agentd/rpc"
	"time"
)

var (
	eventRpcClient *rpc.ConnRPCClient
)

func EventRpcClient() *rpc.ConnRPCClient {
	if eventRpcClient != nil {
		return eventRpcClient
	}
	eventRpcClient = rpc.NewRpcClient(
		cc.Config().RpcAddr,
		time.Duration(cc.Config().Timeout)*time.Second,
	)
	return eventRpcClient
}
