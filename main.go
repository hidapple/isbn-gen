package main

import "os"

var (
	// Name is CLI name
	Name string
	// Version is CLI version
	Version string
	// Revision is build revision
	Revision string
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
