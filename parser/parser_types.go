package parser

const (
	group      uint8 = iota
	bracket    uint8 = iota
	repeat     uint8 = iota
	or         uint8 = iota
	literal    uint8 = iota
	groupUncap uint8 = iota
)

type Token struct {
	val     interface{}
	tokType uint8
}

type RepeatPayload struct {
	min   int
	max   int
	token Token
}

type PContext struct {
	tokens []Token
	index  int
}
