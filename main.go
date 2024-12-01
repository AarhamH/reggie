package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	// Step 1: Take initial user input and store it
	prompt := promptui.Prompt{
		Label: "Enter your first input",
	}

	input1, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	fmt.Printf("Stored input1: %s\n\n", input1)

	// Step 2: Start REPL for multiple inputs
	fmt.Println("Entering REPL mode. Type 'exit' to quit.")

	for {
		// Step 3: Create the prompt for further inputs with cursor support
		prompt = promptui.Prompt{
			Label: fmt.Sprintf("Enter input (you can refer to input1: %s)", input1),
		}

		userInput, err := prompt.Run()
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}

		if strings.ToLower(userInput) == "change" {
			prompt = promptui.Prompt{
				Label: "Enter new input",
			}
			input1, err = prompt.Run()
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				return
			}

			fmt.Printf("Stored input1: %s\n\n", input1)
		}

		// Handle exit condition
		if strings.ToLower(userInput) == "exit" {
			fmt.Println("Exiting REPL.")
			break
		}

		// Optionally process or modify input here
		fmt.Printf("You entered: %s\n", userInput)
	}
}
