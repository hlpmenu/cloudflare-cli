package commands

// Command represents a single command, including its name, description, associated flags, and subcommands.
type Command struct {
	Name        string
	Description string
	Flags       []Flag
	SubCommands []SubCommand // Includes SubCommands to allow for nested commands.
	Run         func()       // Run is a function that will be executed when the command is called.
}

// SubCommand is similar to Command but without SubCommands to avoid infinite nesting.
type SubCommand struct {
	Name        string
	Description string
	Flags       []Flag
	Run         func() // Run is a function that will be executed when the subcommand is called.
}

// Flag represents a single flag for a command, including its name and value.
type Flag struct {
	Name  string
	Value string
}

// Commands holds all available commands.
type Commands struct {
	AvailableCommands []Command
}

// NewCommands initializes and returns a Commands struct without any commands.
// Commands can be added later using the AddCommand method.
func NewCommands() *Commands {
	return &Commands{
		AvailableCommands: []Command{},
	}
}

// AddCommand adds a new command to the Commands struct.
func (c *Commands) Add(cmd Command) {
	c.AvailableCommands = append(c.AvailableCommands, cmd)
}

// AddSubCommand adds a new subcommand to a specific command.
func (cmd *Command) AddSubCommand(subCmd SubCommand) {
	cmd.SubCommands = append(cmd.SubCommands, subCmd)
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
