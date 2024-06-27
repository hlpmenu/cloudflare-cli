package cmd

import (
	"fmt"
	"go-debug/cmd/commands"
	"go-debug/cmd/parse"
	"go-debug/output"
	"log"
)

// Define commands as a package-level variable.
var Cmds *commands.Commands

func Entry() {
	Cmds := commands.Cmds

	Cmds.AvailableCommands = make(map[string]commands.Command)
	Cmds.Get("printhelloworld")
	log.Printf("%v\n", Cmds.AvailableCommands)
	log.Printf("log %v\n", Cmds.AvailableCommands["printhelloworld"])

	example := commands.ExampleCommand()
	Cmds.Add(*example)
	Cmds.Add(*printhelloworld)

	parse.ParseArgs()

}

var printhelloworld = &commands.Command{
	Name:        "printhelloworld",
	Description: "Prints 'Hello, World!'",
	Flags: []commands.Flag{
		commands.Flag{
			Name:  "o",
			Value: "hello",
		},
	},
	Run: helloworld,
}

func helloworld(m map[string]string) {
	log.Printf("Running helloworld command with flags: %v\n", m)
	var txt string

	if m["o"] == "hello" {
		txt = "Hello, World!"
	} else {
		txt = "Goodbye, World!"
	}

	if len(m) == 0 {
		output.Exit("No flags provided")
	}

	for k, v := range m {
		output.Logf("%s: %s\n", k, v)
		log.Printf("%s", k)
		fmt.Println("Key:", k, "Value:", v) // Replace with fmt.Println for debugging

	}

	output.Successf("%s\n", txt)
}
