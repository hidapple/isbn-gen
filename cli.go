package main

import (
	"flag"
	"fmt"
	"io"
	"strconv"
)

const (
	ExitCodeOK    int = iota
	ExitCodeError int = iota
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

	// Bind flag params
	flags.StringVar(&pubcode, "pubcode", "", "Publisher code of ISBN.")
	flags.StringVar(&pubcode, "p", "", "Publisher code of ISBN (Short).")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")
	flags.BoolVar(&version, "v", false, "Print version information and quit (Short).")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	// Validate option: pubcode
	if !isNumber(pubcode) {
		fmt.Fprintf(cli.errStream, "%s: pubcode must be number: %s\n", Name, pubcode)
		return ExitCodeError
	}

	// Show ISBN
	isbn, err := NewIsbn(pubcode)
	if err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err.Error())
		return ExitCodeError
	}
	fmt.Fprintln(cli.outStream, isbn.GetNumber())

	// Succeeded
	return ExitCodeOK
}

func isNumber(pubcode string) bool {
	if len(pubcode) == 0 {
		return true
	}
	if _, err := strconv.Atoi(pubcode); err == nil {
		return true
	}
	return false
}
