package commands

import "log"

// Command represents a single command, including its name, description, associated flags, and subcommands.
type Command struct {
	Name        string
	Description string
	Flags       []Flag
	SubCommands []SubCommand            // Includes SubCommands to allow for nested commands.
	Run         func(map[string]string) // Run is a function that will be executed when the command is called.
}

// SubCommand is similar to Command but without SubCommands to avoid infinite nesting.
type SubCommand struct {
	Name        string
	Description string
	Flags       []Flag
	Run         func(map[string]string) // Run is a function that will be executed when the subcommand is called.
}

// Flag represents a single flag for a command, including its name and value.
type Flag struct {
	Name     string
	Value    string
	HasValue bool // Added a boolean to check if the flag has a value.
}

// Commands holds all available commands.
type Commands struct {
	AvailableCommands map[string]Command
}

// NewCommands initializes a new Commands struct with a pre-populated map of commands.
func NewCommands() *Commands {
	return &Commands{
		AvailableCommands: make(map[string]Command),
	}
}

// Corrected Add method to work with a map
func (c *Commands) Add(cmd Command) {
	c.AvailableCommands[cmd.Name] = cmd
}

// AddSubCommand adds a subcommand to the Command's SubCommands slice.
func (c *Command) AddSubCommand(subCmd SubCommand) {
	c.SubCommands = append(c.SubCommands, subCmd)
}

// Implementing Get method to retrieve a command by name
func (c *Commands) Get(name string) (Command, bool) {
	cmd, exists := c.AvailableCommands[name]
	return cmd, exists
}

// Get the subcommand by name, returns a SubCommand struct and a boolean indicating if the subcommand exists.
func (c *Command) GetSubCommand(name string) (SubCommand, bool) {
	var subCmd SubCommand
	for _, subCmd = range c.SubCommands {
		if subCmd.Name == name {
			return subCmd, true
		}
	}
	return subCmd, false
}

// Usage returns a string describing how to use the command, including its name and description.
func (cmd Command) Usage() string {
	return "program " + cmd.Name + " - " + cmd.Description
}

// Usage returns a string describing how to use the subcommand, including its name and description.
func (subCmd SubCommand) Usage() string {
	return "program [command] " + subCmd.Name + " - " + subCmd.Description
}

// Usage returns a string describing how to use the flag, including its name and value.
func (flag Flag) Usage() string {
	return "--" + flag.Name + " [value] - " + flag.Value
}

// GetFlag returns a Flag struct and a boolean indicating if the flag exists.
func (cmd *Command) GetFlag(name string) (Flag, bool) {
	for _, flag := range cmd.Flags {
		log.Printf("flag: %v\n", flag.Name)
		if flag.Name == name {
			return flag, true
		}
	}
	return Flag{}, false
}

// GetFlag returns a Flag struct and a boolean indicating if the flag exists.
func (cmd *SubCommand) GetFlag(name string) (Flag, bool) {
	for _, flag := range cmd.Flags {
		log.Printf("flag: %v\n", flag.Name)
		if flag.Name == name {
			return flag, true
		}
	}
	return Flag{}, false
}
func FlagExists(m map[string]string, key string) (bool, string) {
	value, exists := m[key]
	return exists, value
}
