package graph

import (
	parser "reggie/parser"
)

func ToGraph(ctx *parser.PContext) *States {
	startState, endState, _ := tokenToFSA(&ctx.Tokens[0])

	for i := 1; i < len(ctx.Tokens); i++ {
		startNext, endNext, _ := tokenToFSA(&ctx.Tokens[i])
		endState.pushTransition(EPSILON, startNext)
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

	endState.pushTransition(EPSILON, end)

	return start
}
