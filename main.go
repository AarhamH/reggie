package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"

	graph "reggie/graph"
	parser "reggie/parser"
)

func main() {
	prompt := promptui.Prompt{
		Label: "Enter a regex formula",
	}

	regExpression, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	ctx := parser.Parse(regExpression)
	g := graph.ToGraph(ctx)

	fmt.Println("Entering REPL mode. Type 'exit()' to quit, or change() to enter new regex formula.")

	for {
		prompt = promptui.Prompt{
			Label: fmt.Sprintf("Enter string to check match (current regex: %s)", regExpression),
		}

		userInput, err := prompt.Run()
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}

		if strings.ToLower(userInput) == "change()" {
			prompt = promptui.Prompt{
				Label:   "Enter new input",
				Default: regExpression,
			}

			regExpression, err = prompt.Run()
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				return
			}

			ctx = parser.Parse(regExpression)
			g = graph.ToGraph(ctx)

			fmt.Printf("Stored input1: %s\n\n", regExpression)
			continue
		}

		if strings.ToLower(userInput) == "exit()" {
			fmt.Println("Exiting REPL.")
			break
		}

		fmt.Printf("You entered: %s\n", userInput)
		result := g.Check(userInput, -1)
		if result {
			fmt.Printf("Match status: true\n")
		} else {
			fmt.Printf("Match status: false\n")
		}
	}
}
