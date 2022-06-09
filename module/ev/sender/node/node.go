package node

import (
	"github.com/GeorgeWang1994/cicada/module/ev/cc"
	"github.com/toolkits/consistent/rings"
	"sort"
)

var (
	JudgeNodeRing *rings.ConsistentHashNodeRing
)

// KeysOfMap 获取map的所有的key
func KeysOfMap(m map[string]string) []string {
	keys := make(sort.StringSlice, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	keys.Sort()
	return keys
}

// InitNodeRings 维护各服务实例节点组成的一个列表。这个列表按照hash算法的结果升序组织。在逻辑上我们把这个列表当作一个环来看待
func InitNodeRings() {
	cfg := cc.Config()
	JudgeNodeRing = rings.NewConsistentHashNodesRing(int32(cfg.Judge.Replicas), KeysOfMap(cfg.Judge.Cluster))
}
