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
		parseBracket(ctx, regInput)
	case '{':
		fmt.Println("Character is a {")
	case '|':
		parseOr(ctx, regInput)
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

func parseBracket(ctx *PContext, regInput string) {
	ctx.index++
	var literals []string
	for regInput[ctx.index] != ']' {
		regChar := regInput[ctx.index]

		if regChar == '-' {
			literalLastIndex := len(literals) - 1

			next := regInput[ctx.index+1]
			prev := literals[literalLastIndex][0]

			literals[literalLastIndex] = fmt.Sprintf("%c%c", prev, next)
			ctx.index++
		} else {
			literals = append(literals, fmt.Sprintf("%c", regChar))
		}

		ctx.index++
	}

	literalsSet := map[uint8]bool{}

	for _, l := range literals {
		for i := l[0]; i <= l[len(l)-1]; i++ {
			literalsSet[i] = true
		}
	}

	ctx.tokens = append(ctx.tokens, Token{
		val:     literalsSet,
		tokType: bracket,
	})
}

func parseOr(ctx *PContext, regInput string) {
	rightContext := &PContext{
		index:  ctx.index,
		tokens: []Token{},
	}

	rightContext.index++

	for rightContext.index < len(regInput) && regInput[rightContext.index] != ')' {
		buildTokens(rightContext, regInput)
		rightContext.index++
	}

	leftToken := Token{
		val:     ctx.tokens,
		tokType: groupUncap,
	}

	rightToken := Token{
		val:     rightContext.tokens,
		tokType: groupUncap,
	}

	ctx.index = rightContext.index
	ctx.tokens = []Token{{
		val:     []Token{leftToken, rightToken},
		tokType: or,
	}}
}
