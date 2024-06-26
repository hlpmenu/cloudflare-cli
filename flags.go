package main

import (
	"fmt"
	"os"
	"strings"
)

type FlagDetail struct {
	ValuesCount int
}

var (
	flagsMap     = make(map[string]FlagDetail)
	dummyCommand = "--sql 'SELECT * FROM users'" // Example command
)

func AddFlag(name string, valuesCount int) {
	flagsMap[name] = FlagDetail{ValuesCount: valuesCount}
}

func ParseCMD(command string) {
	command = dummyCommand
	if len(command) < 4 {
		return
	}

	command = strings.TrimPrefix(command, "cli ")

	exclusiveFlags := []string{"--sql", "--config"} // Example of exclusive flags
	counter := 0
	for _, flag := range exclusiveFlags {
		if strings.Contains(command, flag) {
			counter++
		}
	}
	if counter > 1 {
		fmt.Println("Error: Exclusive flags are set more than once.")
		os.Exit(1)
	}

	args := strings.Split(command, "--")
	for _, arg := range args {
		if arg == "" {
			continue
		}
		split := strings.Fields(arg)
		activeFlag := split[0]
		if detail, exists := flagsMap[activeFlag]; exists && len(split)-1 == detail.ValuesCount {
			activeFlagValue := strings.Join(split[1:], " ")
			fmt.Printf("Flag: %s, Values: %s\n", activeFlag, activeFlagValue)
		}
	}
}
