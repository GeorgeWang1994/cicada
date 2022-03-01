package gg

import (
	"cicada/ev/cc"
	"cicada/ev/rpc"
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
		cc.Config().Rpc.RpcAddr,
		time.Duration(cc.Config().Rpc.Timeout)*time.Second,
	)
	return eventRpcClient
}
