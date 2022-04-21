package function

import (
	"cicada/judge/judge"
	"cicada/pkg/model"
)

type TypeFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, TypeFunction{})
}

func (f TypeFunction) Name() string {
	return "type"
}

func (f TypeFunction) BeforeCompute() error {
	return nil
}

func (f TypeFunction) Compute(strategy model.AlarmStrategy, event interface{}) (isTriggered bool) {
	// 到达指定风险级别才触发告警
	if len(strategy.Type) == 0 {
		return true
	}
	if _, ok := event.(*model.HoneypotEvent); ok {
		if "honeypot_event" == strategy.Type {
			return true
		}
	}
	return false
}

func (f TypeFunction) AfterCompute() error {
	return nil
}
