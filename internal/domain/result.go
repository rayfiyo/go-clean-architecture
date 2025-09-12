package domain

type Result struct {
	Text    *string
	Numbers *[]int
}

func NewTextResult(s string) Result {
	return Result{Text: &s}
}

func NewNumbersResult(ns []int) Result {
	return Result{Numbers: &ns}
}
