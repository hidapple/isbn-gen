package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/hidapple/isbn-gen/isbn"
	"github.com/olekukonko/tablewriter"
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
		code    string
		version bool
		list    bool
	)
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	// Bind flag params
	flags.BoolVar(&version, "version", false, "Print version information and quit")
	flags.BoolVar(&version, "v", false, "Print version information and quit")

	flags.BoolVar(&list, "list", false, "Print supported registration group identifier list")
	flags.BoolVar(&list, "l", false, "Print supported registration group identifier list")

	flags.StringVar(&idGrp, "id", "jp", "Registration group identifier of ISBN")
	flags.StringVar(&idGrp, "id-group", "jp", "Registration group identifier of ISBN")

	flags.StringVar(&code, "code", "", "The code part pregix of ISBN")
	flags.StringVar(&code, "c", "", "The code part pregix of ISBN")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return exitCodeErr
	}

	// Show version
	if version {
		fmt.Fprintf(cli.outStream, "%s version %s (rev: %s)\n", Name, Version, Revision)
		return exitCodeOK
	}

	// Show supported registration group list
	if list {
		cli.printSupportedGroups()
		return exitCodeOK
	}

	// Generate ISBN
	isbn, err := isbn.NewISBN(idGrp, code)
	if err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err.Error())
		return exitCodeErr
	}
	fmt.Fprintln(cli.outStream, isbn.Number())
	return exitCodeOK
}

func (cli *CLI) printSupportedGroups() {
	table := tablewriter.NewWriter(cli.outStream)
	table.SetHeader([]string{"identifying group", "abbreviation", "prefix", "identifier"})
	for _, v := range isbn.Identifiers {
		table.Append([]string{v.GroupName, v.Abbreviation, v.Prefix, v.Identifier})
	}
	table.Render()
}
