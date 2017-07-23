package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
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
	if len(pubcode) > 8 {
		fmt.Fprintf(cli.errStream, "%s: pubcode must be less than 8 digits: %s\n", Name, pubcode)
		return ExitCodeError
	}

	// Generate ISBN
	fmt.Fprintln(cli.outStream, generate(pubcode))

	// Succeeded
	return ExitCodeOK
}

func isNumber(pubcode string) bool {
	if len(pubcode) == 0 {
		return true
	}
	if _, err := strconv.Atoi(pubcode); err != nil {
		return false
	}
	return true
}

func generate(pubcode string) string {
	isbn := generate12digits(pubcode)
	return isbn + calcCheckDigit(isbn)
}

func generate12digits(pubcode string) string {
	const JapanCode = "9784"
	rand.Seed(time.Now().UnixNano())

	isbn := JapanCode + pubcode
	rest := 8 - len(pubcode)
	for i := 0; i < rest; i++ {
		isbn += strconv.Itoa(rand.Intn(10))
	}
	return isbn
}

func calcCheckDigit(isbn12 string) string {
	sum := 0
	for i, v := range strings.Split(isbn12, "") {
		intV, _ := strconv.Atoi(v)
		if i%2 == 0 {
			sum += intV
		} else {
			sum += intV * 3
		}
	}

	calcResult := 10 - (sum % 10)
	if calcResult == 10 {
		return "0"
	} else {
		return strconv.Itoa(calcResult)
	}
}
