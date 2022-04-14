package function

import "cicada/judge/judge"

type SumFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, SumFunction{})
}

func (f SumFunction) Name() string {
	return "sum"
}

func (f SumFunction) BeforeCompute() error {
	return nil
}

func (f SumFunction) Compute(rangeValues []int, operator string, rightValue float64) (isTriggered bool) {
	return true
}

func (f SumFunction) AfterCompute() error {
	return nil
}
