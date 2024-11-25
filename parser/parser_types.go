package parser

const (
	Group      uint8 = iota
	Bracket    uint8 = iota
	Repeat     uint8 = iota
	Or         uint8 = iota
	Literal    uint8 = iota
	GroupUncap uint8 = iota
)

const REPEAT_INDEX = -1

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

// Context Type implementations
func (p *PContext) position() int {
	return p.Index
}

func (p *PContext) increment() int {
	p.Index++
	return p.Index
}

func (p *PContext) incrementTo(newIndex int) {
	p.Index = newIndex
}
