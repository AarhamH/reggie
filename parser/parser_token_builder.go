package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func buildTokens(ctx *PContext, regInput string) {
	regChar := regInput[ctx.Index]

	switch regChar {
	case '(':
		groupPContext := &PContext{
			Index:  ctx.Index,
			Tokens: []Token{},
		}
		parseGroup(groupPContext, regInput)
		ctx.Tokens = append(ctx.Tokens, Token{
			Val:     groupPContext.Tokens,
			TokType: Group,
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
			Val:     regChar,
			TokType: Literal,
		}

		ctx.Tokens = append(ctx.Tokens, token)
	}
}

func parseGroup(ctx *PContext, regInput string) {
	ctx.Index++
	for regInput[ctx.Index] != ')' {
		buildTokens(ctx, regInput)
		ctx.Index++
	}
}

func parseBracket(ctx *PContext, regInput string) {
	ctx.Index++
	var literals []string
	for regInput[ctx.Index] != ']' {
		regChar := regInput[ctx.Index]

		if regChar == '-' {
			literalLastIndex := len(literals) - 1

			next := regInput[ctx.Index+1]
			prev := literals[literalLastIndex][0]

			literals[literalLastIndex] = fmt.Sprintf("%c%c", prev, next)
			ctx.Index++
		} else {
			literals = append(literals, fmt.Sprintf("%c", regChar))
		}

		ctx.Index++
	}

	literalsSet := map[uint8]bool{}

	for _, l := range literals {
		for i := l[0]; i <= l[len(l)-1]; i++ {
			literalsSet[i] = true
		}
	}

	ctx.Tokens = append(ctx.Tokens, Token{
		Val:     literalsSet,
		TokType: Bracket,
	})
}

func parseOr(ctx *PContext, regInput string) {
	rightContext := &PContext{
		Index:  ctx.Index,
		Tokens: []Token{},
	}

	rightContext.Index++

	for rightContext.Index < len(regInput) && regInput[rightContext.Index] != ')' {
		buildTokens(rightContext, regInput)
		rightContext.Index++
	}

	leftToken := Token{
		Val:     ctx.Tokens,
		TokType: GroupUncap,
	}

	rightToken := Token{
		Val:     rightContext.Tokens,
		TokType: GroupUncap,
	}

	ctx.Index = rightContext.Index
	ctx.Tokens = []Token{{
		Val:     []Token{leftToken, rightToken},
		TokType: Or,
	}}
}

func parseRepeating(ctx *PContext, regInput string) {
	regChar := regInput[ctx.Index]

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

	lastTokenIndex := len(ctx.Tokens) - 1
	last := ctx.Tokens[lastTokenIndex]

	ctx.Tokens[lastTokenIndex] = Token{
		Val: RepeatPayload{
			Min:   min,
			Max:   max,
			Token: last,
		},
		TokType: Repeat,
	}
}

func parseRepeatingSpecfic(ctx *PContext, regInput string) {
	startIndex := ctx.Index + 1

	for regInput[ctx.Index] != '}' {
		ctx.Index++
	}

	boundary := regInput[startIndex:ctx.Index]
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

	lastToken := ctx.Tokens[len(ctx.Tokens)-1]
	ctx.Tokens[len(ctx.Tokens)-1] = Token{
		Val: RepeatPayload{
			Min:   min,
			Max:   max,
			Token: lastToken,
		},
		TokType: Repeat,
	}
}
