package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
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

	flags.StringVar(&pubcode, "pubcode", "0", "Publisher code of ISBN")
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
	fmt.Println(generate(pubcode))

	// Succeeded
	return ExitCodeOK
}

func generate(pubcode string) string {
	isbn := generate12digit(pubcode)
	return isbn + strconv.Itoa(calcCheckDigit(isbn))
}

func generate12digit(pubcode string) string {
	const JapanIsbnPrefix = "9784"
	isbn := JapanIsbnPrefix + pubcode
	length := 8 - len(pubcode)
	for i := 0; i < length; i++ {
		isbn += strconv.Itoa(rand.Intn(10))
	}
	return isbn
}

func calcCheckDigit(partOfIsbn string) int {
	sum := 0
	for i, v := range strings.Split(partOfIsbn, "") {
		intV, _ := strconv.Atoi(v)
		if i%2 == 0 {
			sum += intV
		} else if i%2 == 1 {
			sum += intV * 3
		}
	}

	calcResult := 10 - (sum % 10)
	if calcResult == 10 {
		return 0
	} else {
		return calcResult
	}
}
