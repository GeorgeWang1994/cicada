package function

import (
	"github.com/GeorgeWang1994/cicada/module/judge/judge"
	"github.com/GeorgeWang1994/cicada/module/pkg/model"
)

type ThresholdFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, ThresholdFunction{})
}

func (f ThresholdFunction) Name() string {
	return "threshold"
}

func (f ThresholdFunction) BeforeCompute() error {
	return nil
}

func (f ThresholdFunction) Compute(strategy model.AlarmStrategy, event interface{}) (isTriggered bool) {
	if strategy.Value == 0 {
		return true
	}
	if e, ok := event.(*model.SystemEvent); ok {
		if e.Value >= strategy.Value {
			return true
		}
	}
	return false
}

func (f ThresholdFunction) AfterCompute() error {
	return nil
}
