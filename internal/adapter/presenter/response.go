package presenter

import "app/internal/usecase"

// HTTPに依存しない形のレスポンスDTO（handler側でgin.Hにしても良い）
type SuccessJSON struct {
	Text    *string `json:"text,omitempty"`
	Numbers *[]int  `json:"numbers,omitempty"`
}

type ErrorJSON struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func FromUsecase(out usecase.DecodeOutput) SuccessJSON {
	return SuccessJSON{Text: out.Text, Numbers: out.Numbers}
}
