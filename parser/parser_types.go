package parser

const (
	Group      uint8 = iota
	Bracket    uint8 = iota
	Repeat     uint8 = iota
	Or         uint8 = iota
	Literal    uint8 = iota
	GroupUncap uint8 = iota
)

type Token struct {
	Val     interface{}
	TokType uint8
}

type RepeatPayload struct {
	Min   int
	Max   int
	Token Token
}

type PContext struct {
	Tokens []Token
	Index  int
}
