package gg

import (
	"cicada/judge/cc"
	"cicada/judge/rpc"
	"time"
)

var (
	portalRpcClient *rpc.ConnRPCClient
)

func PortalRpcClient() *rpc.ConnRPCClient {
	if portalRpcClient != nil {
		return portalRpcClient
	}
	portalRpcClient = rpc.NewRpcClient(
		cc.Config().Portal.Servers,
		time.Duration(cc.Config().Portal.Timeout),
	)
	return portalRpcClient
}
