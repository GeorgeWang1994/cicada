package queue

import (
	"github.com/GeorgeWang1994/cicada/ev/sender/node"
	"github.com/GeorgeWang1994/cicada/pkg/model"
	log "github.com/sirupsen/logrus"
	nlist "github.com/toolkits/container/list"
)

var (
	JudgeQueues = make(map[string]*nlist.SafeListLimited)
)

// Push2JudgeSendQueue 将数据加入到某个Judge的发送缓存队列, 具体是哪一个Judge 由一致性哈希 决定
func Push2JudgeSendQueue(items []*model.HoneypotEvent) {
	for _, item := range items {
		// 保证每次都能找到到对应的judge节点
		pk := item.PK()
		n, err := node.JudgeNodeRing.GetNode(pk)
		if err != nil {
			log.Printf("get node failed %v", err)
			continue
		}

		Q := JudgeQueues[n]
		_ = Q.PushFront(item)
	}
}
