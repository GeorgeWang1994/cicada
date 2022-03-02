package sender

import (
	"cicada/ev/sender/node"
	"cicada/ev/sender/pool"
	"cicada/ev/sender/task"
	"context"
)

func Start(ctx context.Context) {
	node.InitNodeRings()
	pool.InitConnPools()
	task.StartSendTasks(ctx)
}
