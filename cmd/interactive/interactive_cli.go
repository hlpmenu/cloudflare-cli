package interactive

import (
	"bufio"
	"fmt"
	"go-debug/output"
	"os"
)

// CLI represents the command line interface
type CLI struct {
	reader *bufio.Reader
}

// NewCLI creates a new CLI instance
func NewCLI() *CLI {
	return &CLI{
		reader: bufio.NewReader(os.Stdin),
	}
}

// AskForInput prompts the user with a question and returns their input
func (cli *CLI) AskForInput(question string) string {
	fmt.Println(question)
	input, _ := cli.reader.ReadString('\n')
	return input[:len(input)-1] // Trim the newline character
}

// Run starts the CLI interaction
func (cli *CLI) Run() {
	input := cli.AskForInput("What should I print?")
	fmt.Println("You said:", input)
}

// WaitForEnter prompts the user to press enter to continue
func (cli *CLI) WaitForEnter(prompt string) {
	fmt.Println(prompt)
	cli.reader.ReadString('\n') // Just wait for the user to press enter
}

//
//

// Input prompts the user and returns their input
func (cli *CLI) Input(question string) string {
	fmt.Println(question)
	input, _ := cli.reader.ReadString('\n')
	return input[:len(input)-1] // Trim the newline character
}

// Print displays a message to the user
func (cli *CLI) Print(message string) {
	fmt.Println(message)
}

// Exit terminates the program with an error message using custom output package
func (cli *CLI) Exit(message string) {
	output.Exit(message) // Using custom output.Exit instead of os.Exit
}

// Success indicates successful completion and exits using custom output package
func (cli *CLI) Success(message string) {
	output.Success(message) // Using custom output.Success instead of os.Exit
}

// AutoPrint prints each element of a slice one by one
func (cli *CLI) AutoPrint(messages []string) {
	for _, message := range messages {
		cli.Print(message)
	}
}

// Function
// Execute runs a given function without arguments
func (cli *CLI) Execute(fn func() string) {
	result := fn()
	cli.Print(result)
}
