package main

import (
	"flag"
	"fmt"
	"io"
	"math"
)

const (
	exitCodeOK int = iota
	exitCodeErr
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var (
		idGrp   string
		pubCode string
		repeat  int
		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	// Bind flag params
	flags.StringVar(&idGrp, "id-group", "jp", "Identifying group of ISBN")
	flags.StringVar(&idGrp, "i", "jp", "Identifying group of ISBN (Short)")

	flags.StringVar(&pubCode, "pubcode", "", "Publisher code of ISBN.")
	flags.StringVar(&pubCode, "p", "", "Publisher code of ISBN (Short).")

	flags.IntVar(&repeat, "repeat", 1, "Generate given number of ISBN")
	flags.IntVar(&repeat, "r", 1, "Generate given number of ISBN (Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")
	flags.BoolVar(&version, "v", false, "Print version information and quit (Short).")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return exitCodeErr
	}

	// Show version
	if version {
		fmt.Fprintf(cli.outStream, "%s version %s (rev: %s)\n", Name, Version, Revision)
		return exitCodeOK
	}

	if repeat < 1 {
		fmt.Fprint(cli.errStream, "repeat(r) flag must be a positive number\n")
		return exitCodeErr
	}

	// Validate prefix, pubcode and repeat flag combination
	max := int(math.Pow(10, float64(12-len(PrefixMap[idGrp]+pubCode))))
	if repeat > max {
		fmt.Fprintf(cli.errStream,
			"there are only %d ISBNs starting with %s but repeat option is %s\n",
			max, PrefixMap[idGrp]+pubCode, repeat)
		return exitCodeErr
	}

	// Generate ISBNs
	set := make(map[string]struct{})
	for {
		isbn, err := NewIsbn(idGrp, pubCode)
		if err != nil {
			fmt.Fprintf(cli.errStream, "%v\n", err.Error())
			return exitCodeErr
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
	return exitCodeOK
}
