package gg

import (
	"cicada/agentd/cc"
	"cicada/agentd/rpc"
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
