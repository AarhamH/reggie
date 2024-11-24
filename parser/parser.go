package parser

import (
	"fmt"
)

func parse(regInput string) *PContext {
	ctx := &PContext{
		tokens: []Token{},
		index:  0,
	}

	for ctx.index < len(regInput) {
		buildTokens(ctx, regInput)
		ctx.index++
	}

	return ctx
}

func buildTokens(ctx *PContext, regInput string) {
	regChar := regInput[ctx.index]

	switch regChar {
	case '(':
		fmt.Println("Character is a (")
	case '[':
		fmt.Println("Character is a [")
	case '{':
		fmt.Println("Character is a {")
	case '|':
		fmt.Println("Character is a |")
	case '*', '?', '+':
		fmt.Println("Character is a *, or ? or +")
	default:
		fmt.Println("Character is... something")
	}
}
