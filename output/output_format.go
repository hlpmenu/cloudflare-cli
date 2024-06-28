package output

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func Errorf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Printf("\033[1;31m%s\n%s\n%s\033[0m", "___Error___", msg, "___________")
}

func Warningf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Printf("\033[0;33m%s\n%s\n%s\033[0m", "___Warning___", msg, "____________")
}

func Infof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Printf("___Info___\n%s\n_________", msg)
}

func Logf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Printf("Log: %s", msg)
}

func Successff(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Printf("\033[1;32m%s\n%s\n%s\033[0m", "___Success___", msg, "___________")
}
func Successf(format string, a ...interface{}) {
	// Get the terminal width
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Default to 80 columns if there's an error
		width = 80
	}

	// Create the header and footer lines based on terminal width
	headerFooterLine := strings.Repeat("_", width)

	// Format the message
	msg := fmt.Sprintf(format, a...)
	msgLength := len(msg)

	// Calculate padding
	totalPadding := width - msgLength
	if totalPadding < 0 {
		totalPadding = 0
	}
	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding

	// Create padded message
	paddedMsg := fmt.Sprintf("%s%s%s", strings.Repeat(" ", leftPadding), msg, strings.Repeat(" ", rightPadding))

	// Ensure the message does not exceed the terminal width
	if len(paddedMsg) > width {
		paddedMsg = paddedMsg[:width]
	}

	// Print the message, boxed in by the header and footer
	out := fmt.Sprintf("\033[1;32m%s\n%s\n%s\033[0m", headerFooterLine, paddedMsg, headerFooterLine)
	fmt.Println(out)
}

func Exitf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Fatalf("\033[1;31m%s\n%s\n%s\033[0m", "___Exit___", msg, "_________")
}
