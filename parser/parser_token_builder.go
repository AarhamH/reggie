package parser

import (
	"fmt"
	"strconv"
	"strings"
	regerrors "tinyreg/regerrors"
)

func buildTokens(ctx *PContext, regInput string) *regerrors.RegexError {
	regChar := regInput[ctx.Index]

	switch regChar {
	case '(':
		groupPContext := &PContext{
			Index:  ctx.Index,
			Tokens: []Token{},
		}
		err := parseGroup(groupPContext, regInput)
		if err != nil {
			return err
		}

		token := Token{
			Val:     groupPContext.Tokens,
			TokType: Group,
		}
		ctx.pushToken(token)

	case '[':
		err := parseBracket(ctx, regInput)
		if err != nil {
			return err
		}

	case '{':
		parseRepeatingSpecfic(ctx, regInput)

	case '|':
		err := parseOr(ctx, regInput)
		if err != nil {
			return err
		}

	case '*', '?', '+':
		parseRepeating(ctx, regInput)

	default:
		token := Token{
			Val:     regChar,
			TokType: Literal,
		}

		ctx.pushToken(token)
	}

	return nil
}

func parseGroup(ctx *PContext, regInput string) *regerrors.RegexError {
	if ctx == nil {
		return &regerrors.RegexError{
			Code:    "Context Error",
			Message: "Could not parse group with an empty context",
		}
	}
	ctx.increment()
	for regInput[ctx.Index] != ')' {
		buildTokens(ctx, regInput)
		ctx.increment()
	}

	return nil
}

func parseBracket(ctx *PContext, regInput string) *regerrors.RegexError {
	if ctx == nil {
		return &regerrors.RegexError{
			Code:    "Context error",
			Message: "Trying to parse bracket with a nil context",
		}
	}

	ctx.increment()

	if ctx.Index >= len(regInput) || regInput[ctx.Index] == ']' {
		return &regerrors.RegexError{
			Code:    "Syntax Error",
			Message: "Empty character class or malformed bracket expression",
			Pos:     ctx.Index,
		}
	}

	var literals []string
	for regInput[ctx.Index] != ']' {
		regChar := regInput[ctx.Index]

		if regChar == '-' {
			literalLastIndex := len(literals) - 1
			if len(literals) == 0 || ctx.Index+1 >= len(regInput) {
				return &regerrors.RegexError{
					Code:    "Syntax Error",
					Message: fmt.Sprintf("Invalid range syntax at position %d", ctx.Index),
					Pos:     ctx.Index,
				}
			}

			next := regInput[ctx.Index+1]
			prev := literals[literalLastIndex][0]

			if prev > next {
				return &regerrors.RegexError{
					Code:    "Syntax error",
					Message: fmt.Sprintf("Invalid character range: %c-%c", prev, next),
					Pos:     ctx.Index,
				}
			}

			literals[literalLastIndex] = fmt.Sprintf("%c%c", prev, next)
			ctx.increment()
		} else {
			literals = append(literals, fmt.Sprintf("%c", regChar))
		}

		ctx.increment()
	}

	if ctx.Index >= len(regInput) || regInput[ctx.Index] != ']' {
		return &regerrors.RegexError{
			Code:    "Syntax Error",
			Message: "Unmatched closing bracket ']'",
			Pos:     ctx.Index,
		}
	}

	literalsSet := map[uint8]bool{}
	for _, l := range literals {
		for i := l[0]; i <= l[len(l)-1]; i++ {
			literalsSet[i] = true
		}
	}

	token := Token{
		Val:     literalsSet,
		TokType: Bracket,
	}

	ctx.pushToken(token)

	return nil
}

func parseOr(ctx *PContext, regInput string) *regerrors.RegexError {
	if ctx == nil {
		return &regerrors.RegexError{
			Code:    "Context error",
			Message: "Trying to parse character with a nil context",
		}
	}

	if ctx.Index >= len(regInput) {
		return &regerrors.RegexError{
			Code:    "Context error",
			Message: fmt.Sprintf("Index will reach out of bounds: ctx.Index: %d, length of input: %d", ctx.Index, len(regInput)),
		}
	}

	rightContext := &PContext{
		Index:  ctx.Index,
		Tokens: []Token{},
	}

	rightContext.Index++

	for rightContext.Index < len(regInput) && regInput[rightContext.Index] != ')' {
		err := buildTokens(rightContext, regInput)
		if err != nil {
			return err
		}
		rightContext.Index++
	}

	if rightContext.Index >= len(regInput) || regInput[rightContext.Index] != ')' {
		return &regerrors.RegexError{
			Code:    "Parsing Error Error",
			Message: "Could not reach end paren",
			Pos:     rightContext.Index,
		}
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

	return nil
}

func parseRepeating(ctx *PContext, regInput string) *regerrors.RegexError {
	if ctx == nil {
		return &regerrors.RegexError{
			Code:    "Context error",
			Message: "Trying to parse bracket with a nil context",
		}
	}

	if ctx.Index >= len(regInput) {
		return &regerrors.RegexError{
			Code:    "Context error",
			Message: fmt.Sprintf("Index will reach out of bounds: ctx.Index: %d, length of input: %d", ctx.Index, len(regInput)),
		}
	}

	regChar := regInput[ctx.Index]

	var min int
	var max int

	if regChar == '*' {
		min = 0
		max = REPEAT_INDEX
	} else if regChar == '?' {
		min = 0
		max = 1
	} else {
		min = 1
		max = REPEAT_INDEX
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

	return nil
}

func parseRepeatingSpecfic(ctx *PContext, regInput string) {
	startIndex := ctx.Index + 1

	for regInput[ctx.Index] != '}' {
		ctx.increment()
	}

	boundary := regInput[startIndex:ctx.Index]
	pieces := strings.Split(boundary, ",")

	var min int
	var max int
	if len(pieces) == 1 {
		if value, err := strconv.Atoi(pieces[0]); err != nil {
			panic(err.Error())
		} else {
			min = value
			max = value
		}
	} else if len(pieces) == 2 {
		if value, err := strconv.Atoi(pieces[0]); err != nil {
			panic(err.Error())
		} else {
			min = value
		}

		if pieces[1] == "" {
			max = REPEAT_INDEX
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
