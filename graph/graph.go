package graph

import (
	"tinyreg/parser"
)

func tokenToFSA(t *parser.Token) (*States, *States) {
	start := &States{
		transitions: map[uint8][]*States{},
	}

	end := &States{
		transitions: map[uint8][]*States{},
	}

	switch t.TokType {
	case parser.Literal:
		literalFSA(t, start, end)
	case parser.Or:
	case parser.Bracket:
	case parser.Group, parser.GroupUncap:
	case parser.Repeat:
	default:
		panic("Token type not known")
	}

	return start, end
}

func literalFSA(t *parser.Token, s *States, e *States) {
	char := t.Val.(uint8)
	s.transitions[char] = []*States{e}
}
