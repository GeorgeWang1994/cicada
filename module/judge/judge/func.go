package judge

import (
	"github.com/GeorgeWang1994/cicada/module/pkg/model"
)

type Function interface {
	Name() (name string)
	BeforeCompute() error
	Compute(strategy model.AlarmStrategy, event interface{}) (isTriggered bool)
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
