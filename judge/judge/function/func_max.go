package function

import "cicada/judge/judge"

type MaxFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, MaxFunction{})
}

func (f MaxFunction) Name() string {
	return "max"
}

func (f MaxFunction) BeforeCompute() error {
	return nil
}

func (f MaxFunction) Compute(rangeValues []int, operator string, rightValue float64) (isTriggered bool) {
	return true
}

func (f MaxFunction) AfterCompute() error {
	return nil
}
