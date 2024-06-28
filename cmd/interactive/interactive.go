package interactive

import (
	"go-debug/cmd/commands"
)

var cli = NewCLI()

// main function to run the CLI
func InteractiveTest(m map[string]string) {

	var printboth = false
	for k := range m {
		if k == "2" {
			printboth = true
		}
	}

	cli.Run()
	if printboth {
		cli.WaitForEnter("Press enter to continue")
		cli.Run()
	}

}

var StartInteractive = &commands.Command{
	Name:        "inter",
	Description: "Starts an interactive session",
	Flags: []commands.Flag{
		{
			Name:     "2",
			HasValue: false,
		},
	},
	Run: InteractiveTest,
}
