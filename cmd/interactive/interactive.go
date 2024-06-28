package interactive

import (
	"fmt"
	"go-debug/cmd/commands"
	"log"
)

var cli = NewCLI()

// main function to run the CLI
func InteractiveTest(m map[string]string) {

	if m == nil {
		log.Panicf("m is nil\n")
	}

	var printboth = false
	for k, _ := range m {
		log.Panicf("k: %v\n", k)
		if k == "-2" {
			printboth = true
		}
	}

	cli.Run()

	cli.WaitForInput("Enter a message")
	selection := cli.MultiChoiceQuestion("Run 1 or 2?", []string{"1", "2"})
	if selection == 1 {
		cli.AskForInput("Enter a message")
	} else {
		cli.WaitForInput("ok 2")
	}
	d := fmt.Sprintf("You selected %d", selection)
	cli.Print(d)

	var input string
	for {
		input = cli.AskForInput("Write to exit")
		if input != "" {
			cli.Leave("Exiting")
			break
		}
	}

	if printboth {
		cli.WaitForEnter("Press enter to continue")

	}
	<-cli.quit

}

var StartInteractive = &commands.Command{
	Name:        "inter",
	Description: "Starts an interactive session",
	Flags: []commands.Flag{
		{
			Name:     "-2",
			HasValue: false,
		},
	},
	Run: InteractiveTest,
}
