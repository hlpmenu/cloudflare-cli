package cmd

import "go-debug/cmd/commands"

// Define commands as a package-level variable.
var cmds *commands.Commands

func Entry() {
	cmds = commands.NewCommands()
	example := commands.ExampleCommand()
	cmds.Add(*example)

}
