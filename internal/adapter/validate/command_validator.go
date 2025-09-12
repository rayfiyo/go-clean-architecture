package validate

import (
	"app/internal/domain"
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	ErrEmptyPayload   = errors.New("validation error: empty payload")
	ErrNonASCII       = errors.New("validation error: non-ascii character")
	ErrBadNumberToken = errors.New("validation error: ascii code must be 0..127 integer")
)

type DefaultValidator struct{}

func (v DefaultValidator) Validate(cmd domain.Command) error {
	if strings.TrimSpace(cmd.Payload) == "" {
		return ErrEmptyPayload
	}
	switch cmd.Type {
	case domain.StringToAscii:
		for _, r := range cmd.Payload {
			if r > 127 {
				return ErrNonASCII
			}
		}
	case domain.AsciiToString:
		fields := strings.Fields(cmd.Payload)
		if len(fields) == 0 {
			return ErrEmptyPayload
		}
		for _, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil || n < 0 || n > 127 {
				return ErrBadNumberToken
			}
		}
	default:
		return errors.New("validation error: unknown command type")
	}
	// UTF-8妥当性（StoA側の入力チェック用）
	if !utf8.ValidString(cmd.Payload) {
		return ErrNonASCII
	}
	return nil
}
