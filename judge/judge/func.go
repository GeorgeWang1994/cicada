package judge

import (
	"cicada/judge/judge/function"
	"cicada/pkg/model"
)

type Function interface {
	Name() (name string)
	BeforeCompute() error
	Compute(rangeValues []int, operator string, rightValue float64) (isTriggered bool)
	AfterCompute() error
}

var Functions []Function

func GetFunc(name string) Function {
	for _, f := range Functions {
		if f.Name() == name {
			return f
		}
	}
	return nil
}

// ParseFunc todo: 完善函数解析
func ParseFunc(strategy model.AlarmStrategy) (Function, error) {
	return function.AllFunction{}, nil
}
