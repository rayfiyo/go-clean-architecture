package usecase

import "app/internal/domain"

type DecodeInteractor struct {
	parser    Parser
	validator Validator
	decoder   Decoder
}

func NewDecodeInteractor(p Parser, v Validator, d Decoder) *DecodeInteractor {
	return &DecodeInteractor{parser: p, validator: v, decoder: d}
}

func (uc *DecodeInteractor) Decode(in DecodeInput) (DecodeOutput, error) {
	// 1) Parse
	cmd, err := uc.parser.Parse(in.Raw)
	if err != nil {
		return DecodeOutput{}, err
	}
	// 2) Validate
	if err := uc.validator.Validate(cmd); err != nil {
		return DecodeOutput{}, err
	}
	// 3) Execute
	res, err := uc.decoder.Decode(cmd)
	if err != nil {
		return DecodeOutput{}, err
	}
	// 4) Assemble
	out := DecodeOutput{Text: res.Text, Numbers: res.Numbers}
	return out, nil
}
