package parser

func Parse(regInput string) *PContext {
	ctx := &PContext{
		Tokens: []Token{},
		Index:  0,
	}

	for ctx.position() < len(regInput) {
		buildTokens(ctx, regInput)
		ctx.increment()
	}

	return ctx
}
