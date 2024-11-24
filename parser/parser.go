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
		groupPContext := &PContext{
			index:  ctx.index,
			tokens: []Token{},
		}
		parseGroup(groupPContext, regInput)
		ctx.tokens = append(ctx.tokens, Token{
			val:     groupPContext.tokens,
			tokType: group,
		})
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

func parseGroup(ctx *PContext, regInput string) {
	ctx.index++
	for regInput[ctx.index] != ')' {
		buildTokens(ctx, regInput)
		ctx.index++
	}
}
