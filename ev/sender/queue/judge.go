package queue

import (
	"cicada/ev/sender/node"
	"cicada/pkg/model"
	nlist "github.com/toolkits/container/list"
	"log"
)

var (
	JudgeQueues = make(map[string]*nlist.SafeListLimited)
)

// Push2JudgeSendQueue 将数据加入到某个Judge的发送缓存队列, 具体是哪一个Judge 由一致性哈希 决定
func Push2JudgeSendQueue(items []*model.HoneypotEvent) {
	for _, item := range items {
		pk := item.PK()
		n, err := node.JudgeNodeRing.GetNode(pk)
		if err != nil {
			log.Println("E:", err)
			continue
		}

		Q := JudgeQueues[n]
		_ = Q.PushFront(item)
	}
}
