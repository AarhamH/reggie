package graph

import (
	parser "tinyreg/parser"
)

func ToGraph(ctx *parser.PContext) *States {
	startState, endState := tokenToFSA(&ctx.Tokens[0])

	for i := 1; i < len(ctx.Tokens); i++ {
		startNext, endNext := tokenToFSA(&ctx.Tokens[i])
		endState.Transitions[EPSILON] = append(
			endState.Transitions[EPSILON],
			startNext,
		)
		endState = endNext
	}

	start := &States{
		Transitions: map[uint8][]*States{
			EPSILON: {startState},
		},
		Start: true,
	}
	end := &States{
		Transitions: map[uint8][]*States{},
		End:         true,
	}

	endState.Transitions[EPSILON] = append(
		endState.Transitions[EPSILON],
		end,
	)

	return start
}