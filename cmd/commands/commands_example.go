package commands

import (
	"fmt"
)

var Cmds = NewCommands()

func ExampleCommand() *Command {
	printHelloWorld := func(map[string]string) {
		fmt.Println("Hello, World!")
	}

	helloCommand := Command{
		Name:        "hello",
		Description: "Prints 'Hello, World!'",
		Run:         printHelloWorld,
	}

	return &helloCommand

}
