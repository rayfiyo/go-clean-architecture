package usecase

import "app/internal/domain"

// 入出力DTO（ユースケースはHTTPやJSONを知らない）
type DecodeInput struct {
	Raw string
}

type DecodeOutput struct {
	Text    *string
	Numbers *[]int
}

// ユースケースが依存する抽象（Ports）
type Parser interface {
	Parse(raw string) (domain.Command, error)
}
type Validator interface {
	Validate(cmd domain.Command) error
}
type Decoder interface {
	Decode(cmd domain.Command) (domain.Result, error)
}

// ユースケースI/F
type DecodeInputPort interface {
	Decode(in DecodeInput) (DecodeOutput, error)
}
