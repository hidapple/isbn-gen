package main

import (
	"flag"
	"fmt"
	"io"
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

	flags.BoolVar(&version, "version", false, "Print version information and quit.")
	flags.BoolVar(&version, "v", false, "Print version information and quit (Short).")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintf(cli.errStream, "Parse error: %s\n", err)
		return exitCodeErr
	}

	// Show version
	if version {
		fmt.Fprintf(cli.outStream, "%s version %s (rev: %s)\n", Name, Version, Revision)
		return exitCodeOK
	}

	// Generate ISBNs
	isbn, err := NewISBN(idGrp, pubCode)
	if err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err.Error())
		return exitCodeErr
	}
	fmt.Fprintln(cli.outStream, isbn.Number)

	// Succeeded
	return exitCodeOK
}
