package cmd

import (
	"fmt"
	cfapi "go-debug/cf_api"
	"go-debug/cmd/commands"
	"go-debug/cmd/interactive"
	"go-debug/cmd/parse"
	"go-debug/env"
	"go-debug/output"
	"log"
)

// Define commands as a package-level variable.
var Cmds *commands.Commands

func Entry() {
	env.SetupEnv()
	Cmds := commands.Cmds

	Cmds.AvailableCommands = make(map[string]commands.Command)
	Cmds.Get("printhelloworld")
	log.Printf("%v\n", Cmds.AvailableCommands)
	log.Printf("log %v\n", Cmds.AvailableCommands["printhelloworld"])

	example := commands.ExampleCommand()
	Cmds.Add(*example)
	Cmds.Add(*printhelloworld)
	Cmds.Add(*interactive.StartInteractive)
	cfapi.InitCFApi()

	parse.ParseArgs()

}

var printhelloworld = &commands.Command{
	Name:        "printhelloworld",
	Description: "Prints 'Hello, World!'",
	Flags: []commands.Flag{
		{
			Name:     "-o",
			Value:    "hello",
			HasValue: true,
		},
	},
	Run: helloworld,
}

func helloworld(m map[string]string) {
	log.Printf("Running helloworld command with flags: %v\n", m)
	var txt string

	if m["-o"] == "hello" {
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
