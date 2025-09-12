package domain

type CommandType int

const (
	StringToAscii CommandType = iota
	AsciiToString
)

type Command struct {
	Type    CommandType
	Payload string // StoA: 通常文字列 / AtoS: "65 66 67" のような10進数列（空白区切り）
}
