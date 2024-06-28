package cfapi

import "go-debug/cmd/commands"

var Cmds *commands.Commands

func InitCFApi() {
	// Point to the commands package
	Cmds = commands.Cmds
	// Api groups
	// D1
	loadD1Commands()

}
