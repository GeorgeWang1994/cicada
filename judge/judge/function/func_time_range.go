package function

import (
	"github.com/GeorgeWang1994/cicada/judge/judge"
	"github.com/GeorgeWang1994/cicada/pkg/model"
)

type TimeRangeFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, TimeRangeFunction{})
}

func (f TimeRangeFunction) Name() string {
	return "time_range"
}

func (f TimeRangeFunction) BeforeCompute() error {
	return nil
}

func (f TimeRangeFunction) Compute(strategy model.AlarmStrategy, event interface{}) (isTriggered bool) {
	// 时间在固定范围内才允许触发
	if strategy.StartTime != 0 && strategy.EndTime != 0 {
		// 具体的事件
		if e, ok := event.(*model.HoneypotEvent); ok {
			if strategy.StartTime <= int64(e.StartTime.Second()) &&
				strategy.EndTime >= int64(e.StartTime.Second()) {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func (f TimeRangeFunction) AfterCompute() error {
	return nil
}
