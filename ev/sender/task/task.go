package task

import (
	"cicada/ev/cc"
	"cicada/ev/sender/queue"
	"context"
)

func StartSendTasks(ctx context.Context) {
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
