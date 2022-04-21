package function

import (
	"cicada/judge/judge"
	"cicada/pkg/model"
)

type RiskFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, RiskFunction{})
}

func (f RiskFunction) Name() string {
	return "risk"
}

func (f RiskFunction) BeforeCompute() error {
	return nil
}

func (f RiskFunction) Compute(strategy model.AlarmStrategy, event interface{}) (isTriggered bool) {
	// 到达指定风险级别才触发告警
	if strategy.Risk == 0 {
		return true
	}
	if e, ok := event.(*model.HoneypotEvent); ok {
		if e.RiskLevel >= strategy.Risk {
			return true
		}
	}
	return false
}

func (f RiskFunction) AfterCompute() error {
	return nil
}
