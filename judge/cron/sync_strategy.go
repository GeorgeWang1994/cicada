package cron

import (
	"cicada/judge/cc"
	"cicada/judge/gg"
	"cicada/pkg/model"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

// SynStrategy 从数据库定时同步策略
func SynStrategy() {
	ticker := time.NewTicker(time.Duration(cc.Config().Portal.SyncInterval))
	for {
		select {
		case <-ticker.C:
			var strategiesResponse model.HostStrategiesResponse
			err := gg.PortalRpcClient().Call("GetStrategies", model.NullRpcRequest{}, &strategiesResponse)
			if err != nil {
				log.Println("[ERROR] GetStrategies:", err)
				return
			}

			rebuildStrategyMap(&strategiesResponse)
		}
	}
}

// 重建本地缓存StrategyMap
func rebuildStrategyMap(strategiesResponse *model.HostStrategiesResponse) {
	m := make(map[string][]model.Strategy)
	for _, hs := range strategiesResponse.HostStrategies {
		hostname := hs.Hostname
		for _, strategy := range hs.Strategies {
			key := fmt.Sprintf("%s/%d", hostname, strategy.Id)
			if _, exists := m[key]; exists {
				m[key] = append(m[key], strategy)
			} else {
				m[key] = []model.Strategy{strategy}
			}
		}
	}

	gg.StrategyMap.ReInit(m)
}
