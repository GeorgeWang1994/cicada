package task

import (
	"context"
	"github.com/GeorgeWang1994/cicada/ev/cc"
	"github.com/GeorgeWang1994/cicada/ev/sender/pool"
	"github.com/GeorgeWang1994/cicada/pkg/model"
	"github.com/toolkits/concurrent/semaphore"
	"github.com/toolkits/container/list"
	"time"
)

const DefaultSendTaskSleepInterval = time.Millisecond * 50 // 默认睡眠间隔为50ms

// Judge定时任务, 将 Judge发送缓存中的数据 通过rpc连接池 发送到Judge
func forward2JudgeTask(ctx context.Context, Q *list.SafeListLimited, node string, concurrent int) {
	batch := cc.Config().Judge.Batch // 一次发送,最多batch条数据
	addr := cc.Config().Judge.Cluster[node]
	sema := semaphore.NewSemaphore(concurrent)

	for {
		items := Q.PopBackBy(batch)
		count := len(items)
		if count == 0 {
			time.Sleep(DefaultSendTaskSleepInterval)
			continue
		}

		is := make([]*model.HoneypotEvent, count)
		for i := 0; i < count; i++ {
			is[i] = items[i].(*model.HoneypotEvent)
		}

		//	同步Call + 有限并发 进行发送
		sema.Acquire()
		go func(addr string, judgeItems []*model.HoneypotEvent, count int) {
			defer sema.Release()

			resp := &model.RpcResponse{}
			var err error
			for i := 0; i < 3; i++ { //最多重试3次
				err = pool.JudgeConnPools.Call(ctx, "Judge.ReceiveEvent", judgeItems, resp)
				if err == nil {
					break
				}
				time.Sleep(time.Millisecond * 10)
			}
		}(addr, is, count)
	}
}
