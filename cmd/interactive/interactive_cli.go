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
	quit   chan bool
}

// NewCLI creates a new CLI instance
func NewCLI() *CLI {
	return &CLI{
		reader: bufio.NewReader(os.Stdin),
		quit:   make(chan bool),
	}
}

func (cli *CLI) Run() {
	go func() {
		for {
			select {
			case <-cli.quit:
				return
			default:
				// Keep the CLI alive, waiting for explicit commands or actions
				// CLI will be terminated on sigint
			}
		}
	}()
}

// AskForInput prompts the user with a question and returns their input
func (cli *CLI) AskForInput(question string) string {
	fmt.Println(question)
	input, _ := cli.reader.ReadString('\n')
	return input[:len(input)-1] // Trim the newline character
}

// WaitForEnter prompts the user to press enter to continue
func (cli *CLI) WaitForEnter(prompt string) {
	fmt.Println(prompt)
	cli.reader.ReadString('\n') // Just wait for the user to press enter
}

func (cli *CLI) WaitForInput(message string) {
	fmt.Println(message)
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

func (cli *CLI) MultiChoiceQuestion(question string, choices []string) int {
	fmt.Println(question)
	for i, choice := range choices {
		fmt.Printf("%d: %s\n", i, choice)
	}

	var response int
	for {
		fmt.Print("Enter the number of your choice: ")
		fmt.Scanf("%d", &response)
		if response >= 0 && response < len(choices) {
			break
		}
		fmt.Println("Invalid choice. Please try again.")
	}

	return response
}

// Leave prints a goodbye message and exits the program
func (cli *CLI) Leave(message string) {
	cli.Print(message)
	os.Exit(0)
}

// Function
// Execute runs a given function without arguments
func (cli *CLI) Execute(fn func() string) {
	result := fn()
	cli.Print(result)
}

func ExampleOfBuildingACLI() {
	cli := NewCLI()
	cli.Print("Wecome to the CLI!")

	name := cli.Input("What is your name?")
	cli.Print("Hello, " + name + "!")
	cli.WaitForInput("Press any key to continueS")

	switch cli.MultiChoiceQuestion("What is your favorite color?", []string{"Red", "Green", "Blue"}) {
	case 0:
		cli.Print("Red is a nice color!")
	case 1:
		cli.Print("Green is a nice color!")
	case 2:
		cli.Print("Blue is a nice color!")
	default:
		cli.Print("Invalid choice")
	}

	cli.Leave("Goodbye!")

}
