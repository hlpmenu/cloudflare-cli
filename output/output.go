package output

import (
	"log"
)

func Error(msg string) {
	log.Printf("\033[1;31m%s\n%s\n%s\033[0m", "___Error___", msg, "___________")
}

func Warning(msg string) {
	log.Printf("\033[0;33m%s\n%s\n%s\033[0m", "___Warning___", msg, "____________")
}

func Info(msg string) {
	log.Printf("___Info___\n%s\n_________", msg)
}

func Log(msg string) {
	log.Printf("Log: %s", msg)
}

func Success(msg string) {
	log.Printf("\033[1;32m%s\n%s\n%s\033[0m", "___Success___", msg, "___________")
}

func Exit(msg string) {
	log.Fatalf("\033[1;31m%s\n%s\n%s\033[0m", "___Exit___", msg, "_________")
}
