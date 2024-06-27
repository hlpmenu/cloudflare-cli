package parse

import (
	"go-debug/cmd/commands"
	"go-debug/global"
	"go-debug/output"
	"os"
)

// Example commands
// CMD: cf-cli d1 -d "Description 1" -f1 "Flag 1" -f2 "Flag 2"
// CMD: cf-cli setup env --useconfig
// CMD: cf-cli worker create

type Command struct {
	Name       string
	Flags      []Flag
	SubCommand SubCommand
	Args       []string
	CMD        commands.Command
}
type Flag struct {
	Name  string
	Value string
}
type SubCommand struct {
	Name  string
	flags []Flag
}

func (c Command) CutLast() {
	c.Args = c.Args[:1]
}

func ParseArgs() {

	C := Command{}
	var (
		args = C.Args
		cmd  = C.CMD
	)

	// Get the command line arguments
	args = os.Args[1:] // os.Args[0] is the program name, so we skip it

	if len(args) < 1 {
		output.Error("No command provided")
		output.Exit("Exiting")
		return
	}

	command := args[0]
	var ac = commands.Cmds.AvailableCommands

	// Check if the command is available
	if _, exists := ac[command]; exists {
		output.Infof("Executing command: %s\n", command)

		cmd = ac[command]
		C.CutLast() // Remove the command from the arguments

		for _, arg := range args {
			if arg == "-h" || arg == "--help" {
				output.Info(cmd.Usage())
				return
			} else if arg == "-v" || arg == "--version" {
				output.Info(global.Version)
				return
			} else if regexIsFlag(arg) {
				f, isValid := cmd.GetFlag(arg)
				if !isValid {
					output.Errorf("Unknown flag: %s\n", arg)
					return
				}
				if f.HasValue {
					C.CutLast()
					if len(args) < 1 || regexIsFlag(args[0]) {
						output.Error("No value provided for flag")
						return
					} else {
						C.Flags = append(C.Flags, Flag{Name: f.Name, Value: args[0]})
						C.CutLast()
					}
				}

			}
			if len(args) > 0 {
				break
			}
		}

	} else {
		output.Errorf("Unknown command: %s\n", command)
		// Here you might want to print available commands or usage information
	}

}

func regexIsFlag(s string) bool {
	return len(s) > 1 && s[0] == '-' && (s[1] == '-' || s[1] != '-')
}

func executeCommand(C Command) {
	flags := createFlagMap(C.Flags)
	C.CMD.Run(flags)

}

func createFlagMap(flags []Flag) map[string]string {
	m := make(map[string]string)
	for _, f := range flags {
		m[f.Name] = f.Value
	}
	return m
}
