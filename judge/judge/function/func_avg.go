package function

import "cicada/judge/judge"

type AvgFunction struct {
	judge.Function
}

func init() {
	judge.Functions = append(judge.Functions, AvgFunction{})
}

func (f AvgFunction) Name() string {
	return "avg"
}

func (f AvgFunction) BeforeCompute() error {
	return nil
}

func (f AvgFunction) Compute(rangeValues []int, operator string, rightValue float64) (isTriggered bool) {
	return true
}

func (f AvgFunction) AfterCompute() error {
	return nil
}
