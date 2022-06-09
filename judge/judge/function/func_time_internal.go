package function

import (
	"github.com/GeorgeWang1994/cicada/judge/gg"
	"github.com/GeorgeWang1994/cicada/judge/judge"
	"github.com/GeorgeWang1994/cicada/pkg/model"
	"github.com/garyburd/redigo/redis"
)

type TimeInternalFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, TimeInternalFunction{})
}

func (f TimeInternalFunction) Name() string {
	return "time_internal"
}

func (f TimeInternalFunction) BeforeCompute() error {
	return nil
}

func (f TimeInternalFunction) Compute(strategy model.AlarmStrategy, event interface{}) (isTriggered bool) {
	// 在连续一段时间内达到了阈值，一般使用在数据会发生来回变化的情况下
	// 利用redis缓存过去的值并设置超时时间，下次访问再检查即可
	if strategy.Internal == 0 && strategy.Value == 0 {
		return true
	}
	if strategy.Internal != 0 && strategy.Value != 0 {
		re := gg.RedisConnPool.Get()
		// todo: 完善使用redis lua来写
		if e, ok := event.(*model.SystemEvent); ok {
			s := redis.NewScript(2, "func_time_internal.lua")
			v, err := s.Do(re, strategy.Value, e.Value)
			if err != nil {
				return false
			}
			if v2, ok := v.(int); !ok || v2 == 0 {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func (f TimeInternalFunction) AfterCompute() error {
	return nil
}
