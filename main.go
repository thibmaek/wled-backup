package main

import (
	"fmt"
	"os"
	"strings"
	cmds "thibmaek/wled-export/cmd"
)

type Command interface {
	Name() string
	Run()
}

func main() {
	cmdNames := []string{cmds.ExportCmdName}

	if len(os.Args) < 2 {
		fmt.Printf("Please pass a Command. Options: %s\n", strings.Join(cmdNames, ", "))
		os.Exit(1)
	}

	cmdName := os.Args[1]
	flags := os.Args[2:]

	var cmd Command
	switch cmdName {
	case cmds.ExportCmdName:
		cmd = cmds.NewExportCmd(flags)
	default:
		fmt.Printf("Unknown Command. Options: %s\n", strings.Join(cmdNames, ", "))
		os.Exit(1)
	}

	cmd.Run()
}
