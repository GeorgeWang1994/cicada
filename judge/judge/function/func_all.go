package function

import "cicada/judge/judge"

type AllFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, AllFunction{})
}

func (f AllFunction) Name() string {
	return "all"
}

func (f AllFunction) BeforeCompute() error {
	return nil
}

func (f AllFunction) Compute(rangeValues []int, operator string, rightValue float64) (isTriggered bool) {
	return true
}

func (f AllFunction) AfterCompute() error {
	return nil
}
