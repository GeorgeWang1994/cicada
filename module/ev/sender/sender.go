package sender

import (
	"context"
	"github.com/GeorgeWang1994/cicada/module/ev/sender/node"
	"github.com/GeorgeWang1994/cicada/module/ev/sender/pool"
	"github.com/GeorgeWang1994/cicada/module/ev/sender/task"
)

func Start(ctx context.Context) {
	node.InitNodeRings()
	pool.InitConnPools()
	task.InitSendTasks(ctx)
}
