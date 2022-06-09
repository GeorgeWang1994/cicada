package task

import (
	"context"
	"github.com/GeorgeWang1994/cicada/ev/cc"
	"github.com/GeorgeWang1994/cicada/ev/sender/queue"
)

func InitSendTasks(ctx context.Context) {
	cfg := cc.Config()
	// init semaphore
	judgeConcurrent := cfg.Judge.InitCap

	if judgeConcurrent < 1 {
		judgeConcurrent = 1
	}

	// init send go-routines
	for node := range cfg.Judge.Cluster {
		q := queue.JudgeQueues[node]
		go forward2JudgeTask(ctx, q, node, judgeConcurrent)
	}
}
