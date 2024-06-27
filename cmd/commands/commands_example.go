package commands

import (
	"fmt"
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
