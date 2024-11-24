package parser

import (
	"fmt"
	"strconv"
	"strings"
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
		parseRepeatingSpecfic(ctx, regInput)
	case '|':
		parseOr(ctx, regInput)
	case '*', '?', '+':
		parseRepeating(ctx, regInput)
	default:
		token := Token{
			val:     regChar,
			tokType: literal,
		}

		ctx.tokens = append(ctx.tokens, token)
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

func parseRepeating(ctx *PContext, regInput string) {
	regChar := regInput[ctx.index]

	var min int
	var max int

	if regChar == '*' {
		min = 0
		max = -1
	} else if regChar == '?' {
		min = 0
		max = 1
	} else {
		min = 1
		max = -1
	}

	lastTokenIndex := len(ctx.tokens) - 1
	last := ctx.tokens[lastTokenIndex]

	ctx.tokens[lastTokenIndex] = Token{
		val: RepeatPayload{
			min:   min,
			max:   max,
			token: last,
		},
		tokType: repeat,
	}
}

func parseRepeatingSpecfic(ctx *PContext, regInput string) {
	startIndex := ctx.index + 1

	for regInput[ctx.index] != '}' {
		ctx.index++
	}

	boundary := regInput[startIndex:ctx.index]
	pieces := strings.Split(boundary, ",")

	var min int
	var max int
	if len(pieces) == 1 { // <3>
		if value, err := strconv.Atoi(pieces[0]); err != nil {
			panic(err.Error())
		} else {
			min = value
			max = value
		}
	} else if len(pieces) == 2 { // <4>
		if value, err := strconv.Atoi(pieces[0]); err != nil {
			panic(err.Error())
		} else {
			min = value
		}

		if pieces[1] == "" {
			max = -1
		} else if value, err := strconv.Atoi(pieces[1]); err != nil {
			panic(err.Error())
		} else {
			max = value
		}
	} else {
		panic(fmt.Sprintf("There must be either 1 or 2 values specified for the quantifier: provided '%s'", boundary))
	}

	lastToken := ctx.tokens[len(ctx.tokens)-1]
	ctx.tokens[len(ctx.tokens)-1] = Token{
		val: RepeatPayload{
			min:   min,
			max:   max,
			token: lastToken,
		},
		tokType: repeat,
	}
}
