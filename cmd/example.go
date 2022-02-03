package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	exampleCmdName = "example"
)

type ExampleCmd struct {
	name    string
	verbose bool
}

func NewExampleCmd(flags []string) *ExampleCmd {
	cmd := &ExampleCmd{name: exampleCmdName}

	fs := flag.NewFlagSet(exampleCmdName, flag.ExitOnError)
	fs.BoolVar(&cmd.verbose, "verbose", false, "Show verbose output")

	err := fs.Parse(flags)
	if err != nil {
		fs.PrintDefaults()
		os.Exit(1)
	}

	return cmd
}

func (cmd *ExampleCmd) Name() string {
	return cmd.name
}

func (cmd *ExampleCmd) Run() {}
