package main

import (
	"flag"
	"fmt"
	"io"
	"math"
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
		repeat  int
		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	// Bind flag params
	flags.StringVar(&pubcode, "pubcode", "", "Publisher code of ISBN.")
	flags.StringVar(&pubcode, "p", "", "Publisher code of ISBN (Short).")

	flags.IntVar(&repeat, "repeat", 1, "Generate specified number of ISBN")
	flags.IntVar(&repeat, "r", 1, "Generate specified number of ISBN")

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

	if repeat < 1 {
		fmt.Fprint(cli.errStream, "repeat(r) flag must be a positive number\n")
		return ExitCodeError
	}

	// Validate pubcode and repeat flag combination
	max := int(math.Pow(10, float64(8-len(pubcode))))
	if repeat > max {
		fmt.Fprintf(cli.errStream, "There are only %d ISBNs that can be generated with pubcode:%s\n", max, pubcode)
		return ExitCodeError
	}

	// Generate ISBNs
	set := make(map[string]struct{})
	for {
		isbn, err := NewIsbn(pubcode)
		if err != nil {
			fmt.Fprintf(cli.errStream, "%v\n", err.Error())
			return ExitCodeError
		}
		set[isbn.Number] = struct{}{}

		if len(set) == repeat {
			break
		}
	}

	// Show ISBN
	var output string
	for k, _ := range set {
		output += k + "\n"
	}
	fmt.Fprint(cli.outStream, output)

	// Succeeded
	return ExitCodeOK
}
