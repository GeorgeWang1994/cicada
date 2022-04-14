package function

import "cicada/judge/judge"

type MinFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, MinFunction{})
}

func (f MinFunction) Name() string {
	return "min"
}

func (f MinFunction) BeforeCompute() error {
	return nil
}

func (f MinFunction) Compute(rangeValues []int, operator string, rightValue float64) (isTriggered bool) {
	return true
}

func (f MinFunction) AfterCompute() error {
	return nil
}
