package parse

import (
	"go-debug/cmd/commands"
	"go-debug/global"
	"go-debug/output"
	"go-debug/output/timer"
	"log"
	"os"
	"strings"
)

// Example commands
// CMD: cf-cli d1 -d "Description 1" -f1 "Flag 1" -f2 "Flag 2"
// CMD: cf-cli setup env --useconfig
// CMD: cf-cli worker create

type Command struct {
	Name         string
	Flags        []Flag
	SubCommand   commands.SubCommand
	IsSubCommand bool
	Args         Args
	CMD          commands.Command
}
type Flag struct {
	Name  string
	Value string
}

type Args []string

func (args *Args) CutLast() {
	var new Args = make(Args, len(*args)) // Initialize 'new' with the same length as 'args'
	copy(new, *args)                      // Copy the contents of 'args' into 'new'
	if len(*args) > 0 {
		new = new[1:] // Modify 'new' by removing the first element
		*args = new   // Update 'args' with the value of 'new'
	}
}

func ParseArgs() {
	timer.Time("ParseArgs")()
	C := &Command{}
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
		log.Print("command: ", command)
		cmd = ac[command]

		args.CutLast()

		log.Print("args: ", args)
		loopCount := 0
		for _, arg := range args {
			loopCount++
			if arg == "-h" || arg == "--help" {
				output.Info(cmd.Usage())
				return
			} else if arg == "-v" || arg == "--version" {
				output.Info(global.Version)
				return
			}
			if regexIsFlag(arg) {
				log.Printf("arg in if: %s\n", arg)
				var (
					isValid bool
					f       commands.Flag
				)
				if !C.IsSubCommand {
					f, isValid = cmd.GetFlag(arg)
				} else if C.IsSubCommand {
					f, isValid = C.SubCommand.GetFlag(arg)
				}
				log.Printf("f: %v\n, %v, %v", f, isValid, arg)
				if !isValid {
					output.Errorf("Unknown flag: %s\n", arg)
					return
				}
				if f.HasValue {
					args.CutLast()
					if len(args) < 1 || regexIsFlag(args[0]) {
						output.Error("No value provided for flag")
						return
					} else {
						C.Flags = append(C.Flags, Flag{Name: f.Name, Value: args[0]})
						args.CutLast()
					}
				}
			} else {
				log.Printf("arg in else: %s\n", arg)
				log.Printf("args in else: %v\n", args)
				// Check if the argument is a subcommand
				subCmd, isValid := cmd.GetSubCommand(arg)
				if isValid {
					C.IsSubCommand = true
					C.SubCommand = subCmd
					args.CutLast() // Remove the subcommand from args for further processing
				} else {
					output.Errorf("Unknown command or argument: %s\n", arg)
					return
				}
			}
			if len(args) < 1 {
				log.Printf("Break: %s\n", args)
				break
			}

		}
		log.Printf("count: %v\n", loopCount)

		log.Printf(command)
		C.CMD = cmd
		executeCommand(C)

	} else {
		output.Errorf("Unknown command: %s\n", command)
		// Here you might want to print available commands or usage information
	}

}

func regexIsFlag(s string) bool {
	//return len(s) > 1 && s[0] == '-' && (s[1] == '-' || s[1] != '-')
	return strings.HasPrefix(s, "-")
}

func executeCommand(C *Command) {

	flags := createFlagMap(C.Flags)
	if flags == nil { // Added to show you that the flags map is NOT nil.
		output.Error("Error creating flags map")
		output.Exit("Exiting...")
	}

	if C.IsSubCommand {
		if C.SubCommand.Run == nil {
			output.Error("Internal error: Subcommand has no run function")
		}
		C.SubCommand.Run(flags)
	} else if C.CMD.Run != nil {
		C.CMD.Run(flags)
	} else {
		output.Error("Internal error: Command has no run function")
		output.Exit("Exiting...")
	}

}

func createFlagMap(flags []Flag) map[string]string {
	m := make(map[string]string)
	for _, f := range flags {
		m[f.Name] = f.Value
	}
	return m
}
