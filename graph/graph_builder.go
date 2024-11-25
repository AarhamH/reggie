package graph

import parser "tinyreg/parser"

func tokenToFSA(t *parser.Token) (*States, *States) {
	start := &States{
		Transitions: map[uint8][]*States{},
	}

	end := &States{
		Transitions: map[uint8][]*States{},
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
		repeatFSA(t, start, end)
	default:
		panic("Token type not known")
	}

	return start, end
}

func literalFSA(t *parser.Token, s *States, e *States) {
	char := t.Val.(uint8)
	s.Transitions[char] = []*States{e}
}

func orFSA(t *parser.Token, s *States, e *States) {
	vals := t.Val.([]parser.Token)
	left := vals[0]
	right := vals[1]

	s1, e1 := tokenToFSA(&left)
	s2, e2 := tokenToFSA(&right)

	s.Transitions[EPSILON] = []*States{s1, s2}
	e1.Transitions[EPSILON] = []*States{e}
	e2.Transitions[EPSILON] = []*States{e}
}

func bracketFSA(t *parser.Token, s *States, e *States) {
	literals := t.Val.(map[uint8]bool)

	for l := range literals {
		s.Transitions[l] = []*States{e}
	}
}

func groupFSA(t *parser.Token, s *States, e *States) {
	tokens := t.Val.([]parser.Token)
	s, e = tokenToFSA(&tokens[0])

	for i := 1; i < len(tokens); i++ {
		ts, te := tokenToFSA(&tokens[i])
		e.pushTransition(EPSILON, ts)
		e = te
	}
}

func repeatFSA(t *parser.Token, s *States, e *States) {
	p := t.Val.(parser.RepeatPayload)

	if p.Min == 0 {
		s.Transitions[EPSILON] = []*States{e}
	}

	var copyCount int

	if p.Max == parser.REPEAT_INDEX {
		if p.Min == 0 {
			copyCount = 1
		} else {
			copyCount = p.Min
		}
	} else {
		copyCount = p.Max
	}

	from, to := tokenToFSA(&p.Token)
	s.pushTransition(EPSILON, from)

	for i := 2; i <= copyCount; i++ {
		s, e := tokenToFSA(&p.Token)

		to.pushTransition(EPSILON, s)

		from = s
		to = e

		if i > p.Min {
			s.pushTransition(EPSILON, e)
		}
	}
	to.pushTransition(EPSILON, e)

	if p.Max == parser.REPEAT_INDEX {
		e.pushTransition(EPSILON, from)
	}
}
