package output

import (
	"fmt"
	"log"
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

func Successf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Printf("\033[1;32m%s\n%s\n%s\033[0m", "___Success___", msg, "___________")
}

func Exitf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Fatalf("\033[1;31m%s\n%s\n%s\033[0m", "___Exit___", msg, "_________")
}
