package graph

import (
	parser "tinyreg/parser"
)

const EPSILON uint8 = 0

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
		orFSA(t, start, end)
	case parser.Bracket:
		bracketFSA(t, start, end)
	case parser.Group, parser.GroupUncap:
		groupFSA(t, start, end)
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

func orFSA(t *parser.Token, s *States, e *States) {
	vals := t.Val.([]parser.Token)
	left := vals[0]
	right := vals[1]

	s1, e1 := tokenToFSA(&left)
	s2, e2 := tokenToFSA(&right)

	s.transitions[EPSILON] = []*States{s1, s2}
	e1.transitions[EPSILON] = []*States{e}
	e2.transitions[EPSILON] = []*States{e}
}

func bracketFSA(t *parser.Token, s *States, e *States) {
	literals := t.Val.(map[uint8]bool)

	for l := range literals {
		s.transitions[l] = []*States{e}
	}
}

func groupFSA(t *parser.Token, s *States, e *States) {
	tokens := t.Val.([]parser.Token)
	s, e = tokenToFSA(&tokens[0])

	for i := 1; i < len(tokens); i++ {
		ts, te := tokenToFSA(&tokens[i])
		e.transitions[EPSILON] = append(e.transitions[EPSILON], ts)
		e = te
	}
}
