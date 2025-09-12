package parse

import (
	"app/internal/domain"
	"errors"
	"regexp"
	"strings"
)

// かなり単純化：
// - 「数字と空白のみ」→ AtoS
// - それ以外が含まれる → StoA
// - 空や混在不明瞭 → エラー
var digitsSpaces = regexp.MustCompile(`^[\s0-9]+$`)

type WhitespaceParser struct{}

var ErrParse = errors.New("parse error: could not determine command type")

func (p WhitespaceParser) Parse(raw string) (domain.Command, error) {
	trim := strings.TrimSpace(raw)
	if trim == "" {
		return domain.Command{}, ErrParse
	}
	if digitsSpaces.MatchString(trim) {
		// 数字列：AtoS
		return domain.Command{Type: domain.AsciiToString, Payload: trim}, nil
	}
	// それ以外：StoA（記号や英字を含む通常テキスト）
	return domain.Command{Type: domain.StringToAscii, Payload: trim}, nil
}
