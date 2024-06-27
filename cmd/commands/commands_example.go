package commands

import (
	"fmt"
	"go-debug/output"
)

func ExampleCommand() *Command {
	printHelloWorld := func() {
		fmt.Println("Hello, World!")
	}

	helloCommand := Command{
		Name:        "hello",
		Description: "Prints 'Hello, World!'",
		Run:         printHelloWorld,
	}

	return &helloCommand

}

func ExampleSubCommand() *SubCommand {
	printHelloWorld := func() {
		fmt.Println("Hello, City!")
	}

	helloCommand := SubCommand{
		Name:        "hello",
		Description: "Prints 'Hello, World!'",
		Flags: []Flag{
			Flag{
				Name: "city",
			},
			Flag{
				Name:  "country",
				Value: "USA",
			},
		},
		Run: printHelloWorld,
	}

	return &helloCommand

}

func TestTheFlags(value string) {
	fmt.Println("The value is: ", value)
}

func AFuncThatDefinesACommand() {
	cmd := Command{
		Name:        "hello-world",
		Description: "Prints 'Hello, World!'",
		Run:         printHello(),
	}
}

func printHello(s string) {
	output.Success()("Hello, ", s)
}
