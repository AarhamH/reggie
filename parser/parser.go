package parser

func parse(regInput string) {
	ctx := &PContext{
		tokens: []Token{},
		index:  0,
	}

	for ctx.index < len(regInput) {
		// process the regex

		ctx.index++
	}
}
