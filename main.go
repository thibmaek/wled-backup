package main

import (
	"fmt"
	"os"
	"thibmaek/go-simple-cli/cmd"
)

type Command interface {
	Name() string
	Run()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please pass a Command. Options:")
		os.Exit(1)
	}

	executable := os.Args[0]
	cmdName := os.Args[1]
	flags := os.Args[2:]

	var cmd Command
	switch cmdName {
	case ExampleCmdName:
		cmd = NewExampleCommand(flags)
	default:
		fmt.Println("Unknown Command. Options:") 
		os.Exit(1)
	}

	cmd.Run()
}
