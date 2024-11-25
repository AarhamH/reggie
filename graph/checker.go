package graph

const (
	startOfText uint8 = 1
	endOfText   uint8 = 2
)

func getChar(input string, pos int) uint8 {
	if pos >= len(input) {
		return endOfText
	}

	if pos < 0 {
		return startOfText
	}

	return input[pos]
}

func (s *States) Check(input string, pos int) bool {
	ch := getChar(input, pos)

	if ch == endOfText && s.End {
		return true
	}

	if states := s.Transitions[ch]; len(states) > 0 {
		nextState := states[0]
		if nextState.Check(input, pos+1) {
			return true
		}
	}

	for _, state := range s.Transitions[EPSILON] {
		if state.Check(input, pos) {
			return true
		}

		if ch == startOfText && state.Check(input, pos+1) {
			return true
		}
	}

	return false
}
