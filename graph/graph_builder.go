package graph

import (
	parser "reggie/parser"
	regerrors "reggie/regerrors"
)

func tokenToFSA(t *parser.Token) (*States, *States, *regerrors.RegexError) {
	start := &States{
		Transitions: map[uint8][]*States{},
	}

	end := &States{
		Transitions: map[uint8][]*States{},
	}

	switch t.TokType {
	case parser.Literal:
		err := literalFSA(t, start, end)
		if err != nil {
			return nil, nil, err
		}
	case parser.Or:
		err := orFSA(t, start, end)
		if err != nil {
			return nil, nil, err
		}

	case parser.Bracket:
		err := bracketFSA(t, start, end)
		if err != nil {
			return nil, nil, err
		}

	case parser.Group, parser.GroupUncap:
		err := groupFSA(t, start, end)
		if err != nil {
			return nil, nil, err
		}
	case parser.Repeat:
		err := repeatFSA(t, start, end)
		if err != nil {
			return nil, nil, err
		}

	default:
		return nil, nil, &regerrors.RegexError{
			Code:    "Graph Error",
			Message: "Token type not found",
		}
	}

	return start, end, nil
}

func literalFSA(t *parser.Token, s *States, e *States) *regerrors.RegexError {
	if t == nil || s == nil || e == nil {
		return &regerrors.RegexError{
			Code:    "Graph Error",
			Message: "Cannot build graph with nil token, start, and/or end states",
		}
	}
	char := t.Val.(uint8)
	s.Transitions[char] = []*States{e}

	return nil
}

func orFSA(t *parser.Token, s *States, e *States) *regerrors.RegexError {
	if t == nil || s == nil || e == nil {
		return &regerrors.RegexError{
			Code:    "Graph Error",
			Message: "Cannot build graph with nil token, start, and/or end states",
		}
	}

	vals := t.Val.([]parser.Token)
	left := vals[0]
	right := vals[1]

	s1, e1, err := tokenToFSA(&left)
	if err != nil {
		return err
	}

	s2, e2, err := tokenToFSA(&right)
	if err != nil {
		return err
	}

	s.Transitions[EPSILON] = []*States{s1, s2}
	e1.Transitions[EPSILON] = []*States{e}
	e2.Transitions[EPSILON] = []*States{e}

	return nil
}

func bracketFSA(t *parser.Token, s *States, e *States) *regerrors.RegexError {
	if t == nil || s == nil || e == nil {
		return &regerrors.RegexError{
			Code:    "Graph Error",
			Message: "Cannot build graph with nil token, start, and/or end states",
		}
	}

	literals := t.Val.(map[uint8]bool)
	if literals == nil {
		return &regerrors.RegexError{
			Code:    "Graph Error",
			Message: "Literals set is nil",
		}
	}

	for l := range literals {
		s.Transitions[l] = []*States{e}
	}

	return nil
}

func groupFSA(t *parser.Token, s *States, e *States) *regerrors.RegexError {
	if t == nil || s == nil || e == nil {
		return &regerrors.RegexError{
			Code:    "Graph Error",
			Message: "Cannot build graph with nil token, start, and/or end states",
		}
	}

	tokens := t.Val.([]parser.Token)
	_, e, err := tokenToFSA(&tokens[0])
	if err != nil {
		return err
	}

	for i := 1; i < len(tokens); i++ {
		ts, te, _ := tokenToFSA(&tokens[i])
		e.pushTransition(EPSILON, ts)
		e = te
	}

	return nil
}

func repeatFSA(t *parser.Token, s *States, e *States) *regerrors.RegexError {
	if t == nil || s == nil || e == nil {
		return &regerrors.RegexError{
			Code:    "Graph Error",
			Message: "Cannot build graph with nil token, start, and/or end states",
		}
	}
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

	from, to, _ := tokenToFSA(&p.Token)
	s.pushTransition(EPSILON, from)

	for i := 2; i <= copyCount; i++ {
		s, e, _ := tokenToFSA(&p.Token)

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

	return nil
}
