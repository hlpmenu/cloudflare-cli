package cmd

import (
	"fmt"
	cfapi "go-debug/cf_api"
	"go-debug/cmd/commands"
	"go-debug/cmd/interactive"
	"go-debug/cmd/parse"
	"go-debug/env"
	setenv "go-debug/env/set_env"
	"go-debug/output"
	"log"
	"os"
	"strings"
)

// Define commands as a package-level variable.
var Cmds *commands.Commands

func Entry() {
	env.SetupEnv()
	Cmds := commands.Cmds

	Cmds.AvailableCommands = make(map[string]commands.Command)
	Cmds.Get("printhelloworld")

	example := commands.ExampleCommand()
	Cmds.Add(*example)
	Cmds.Add(*printhelloworld)
	Cmds.Add(*interactive.StartInteractive)
	Cmds.Add(*printenvcommand)
	cfapi.InitCFApi()
	Cmds.Add(*setenv.ENVCOMMAND)
	setenv.ENVCOMMAND.AddSubCommand(*setenv.SETENVCOMMAND)

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

var printenvcommand = &commands.Command{
	Name:        "printenv",
	Description: "Prints the environment variables",
	Flags: []commands.Flag{
		{
			Name:     "-cf",
			HasValue: false,
		},
	},
	Run: printenv,
}

func printenv(m map[string]string) {
	cfarray := []string{env.CLOUDFLARE_ACCOUNT_EMAIL, env.CLOUDFLARE_API_KEY, env.DB_ID, env.PAGES_ID, env.WORKERS_ID}

	cfonly, _ := commands.FlagExists(m, "-cf")
	if cfonly {
		for _, v := range cfarray {
			log.Printf("%s\n", v)
		}
		return
	} else {
		e := make(map[string]string)
		env := os.Environ()
		for _, i := range env {
			kv := strings.Split(i, "=")
			e[kv[0]] = kv[1]
			log.Printf("%s: %s\n", kv[0], kv[1])
		}
	}

}
