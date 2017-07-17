package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var (
		pubcode string
		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&pubcode, "-pubcode", "0", "Publisher code of ISBN")
	flags.StringVar(&pubcode, "p", "0", "Publisher code of ISBN (Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

  // Generate ISBN
  isbn := generate(pubcode)
  fmt.Print(isbn);

	return ExitCodeOK
}

func generate(pubcode string) string {
  const JapanIsbnPrefix = "9784"
  return JapanIsbnPrefix + pubcode
}
