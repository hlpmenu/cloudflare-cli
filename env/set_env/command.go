package setenv

import (
	"go-debug/cmd/commands"
	output "go-debug/output"
	"os"
	"strings"
)

type flagsMap map[string]string

type ENV_COMMAND struct {
	CMD   string
	Flags map[string]string
}

var ENVCOMMAND = &commands.Command{
	Name:        "env",
	Description: "Manage environment variables",
	Flags: []commands.Flag{
		{
			Name:     "-remote",
			HasValue: false,
		},
		{
			Name:     "-set",
			HasValue: true,
		},
		// Add Id flag
	},
	Run: func(m map[string]string) {
		ENV_ENTRYPOINT(&ENV_COMMAND{CMD: "create", Flags: m})
	},
}
var SETENVCOMMAND = &commands.SubCommand{
	Name:        "set",
	Description: "Set environment variables",
	Flags:       []commands.Flag{},
	Run: func(m map[string]string) {
		ENV_ENTRYPOINT(&ENV_COMMAND{CMD: "set", Flags: m})
	},
}

func ENV_ENTRYPOINT(c *ENV_COMMAND) {

	flags := c.Flags
	CMD := c.CMD

	switch CMD {
	case "create":
		output.Error("Not implemented yet")
	case "set":
		Setup(flags)
	case "env":

	default:
		output.Error("Not implemented yet")
	}

	exists, s := commands.FlagExists(flags, "-set")
	if exists && len(s) > 0 {
		pair := strings.Split(s, "=")
		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])
		os.Setenv(key, value)
	}

}

func Setup(m flagsMap) {

	SetEnv()

}
