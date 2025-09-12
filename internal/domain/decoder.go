package domain

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	ErrNonASCII   = errors.New("non-ascii character detected")
	ErrOutOfRange = errors.New("ascii code out of range")
)

type Decoder interface {
	Decode(Command) (Result, error)
}

type DefaultDecoder struct{}

func (d DefaultDecoder) Decode(c Command) (Result, error) {
	switch c.Type {
	case StringToAscii:
		// 各文字をASCII 10進へ
		var nums []int
		for _, r := range c.Payload {
			if r > 127 { // ASCII外
				return Result{}, ErrNonASCII
			}
			nums = append(nums, int(r))
		}
		return NewNumbersResult(nums), nil

	case AsciiToString:
		fields := strings.Fields(c.Payload)
		var b strings.Builder
		for _, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil || n < 0 || n > 127 {
				return Result{}, ErrOutOfRange
			}
			b.WriteByte(byte(n))
		}
		s := b.String()
		// 生成結果はASCIIなので必ずUTF-8の1バイト並び（ただし一応検査）
		if !utf8.ValidString(s) {
			return Result{}, ErrOutOfRange
		}
		return NewTextResult(s), nil

	default:
		return Result{}, errors.New("unknown command type")
	}
}
