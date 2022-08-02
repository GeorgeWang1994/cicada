package cron

import (
	"github.com/GeorgeWang1994/cicada/module/judge/cc"
	"github.com/GeorgeWang1994/cicada/module/judge/gg"
	"github.com/GeorgeWang1994/cicada/module/pkg/model"
	log "github.com/sirupsen/logrus"
	"time"
)

// SyncAlarmStrategy 从数据库定时同步告警策略
func SyncAlarmStrategy() {
	ticker := time.NewTicker(time.Duration(cc.Config().Portal.SyncInterval))
	for {
		select {
		case <-ticker.C:
			var strategiesResponse model.StrategiesResponse
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
func rebuildStrategyMap(strategiesResponse *model.StrategiesResponse) {
	m := make([]model.AlarmStrategy, 0)
	for _, strategy := range strategiesResponse.Strategies {
		m = append(m, strategy)
	}
	gg.AlarmStrategy.ReInit(m)
}

// SyncSubscribeStrategy 从数据库定时同步订阅策略
func SyncSubscribeStrategy() {
	ticker := time.NewTicker(time.Duration(cc.Config().Portal.SyncInterval))
	for {
		select {
		case <-ticker.C:
			var strategiesResponse model.StrategiesResponse
			err := gg.PortalRpcClient().Call("GetStrategies", model.NullRpcRequest{}, &strategiesResponse)
			if err != nil {
				log.Println("[ERROR] GetStrategies:", err)
				return
			}

			rebuildStrategyMap(&strategiesResponse)
		}
	}
}
