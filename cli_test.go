package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	Name = "isbn-gen"
	Version = "v1.0"
	Revision = "abcdef"
	args := strings.Split("./isbn-gen --version", " ")

	status := cli.Run(args)
	if status != exitCodeOK {
		t.Fatalf("Expected exit code is %d but was %d", exitCodeOK, status)
	}

	expected := fmt.Sprintf("isbn-gen version %s (rev: %s)", Version, Revision)
	if !strings.Contains(outStream.String(), expected) {
		t.Fatalf("Expected output contain %q but was %q", expected, errStream.String())
	}
}

func TestRun_listFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen --list", " ")

	status := cli.Run(args)

	// exitCode should be 0
	if status != exitCodeOK {
		t.Fatalf("Expected exit code is %d but was %d", exitCodeOK, status)
	}

	// Supported identifier group table should be shown
	golden := filepath.Join("testdata", t.Name()+".golden")
	expected, _ := ioutil.ReadFile(golden)
	if outStream.String() != string(expected) {
		t.Fatalf("Expected output is %s but was %s", expected, outStream.String())
	}
}

func TestRun_idGroupFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen --id-group en", " ")

	status := cli.Run(args)

	// exitCode should be 0
	if status != exitCodeOK {
		t.Fatalf("Expected exit code is %d but was %d", exitCodeOK, status)
	}

	// Output ISBN should contain 9784(en1 prefix)
	expected := "9780"
	if !strings.Contains(outStream.String(), expected) {
		t.Fatalf("Expected output contain %q but was %q", expected, outStream.String())
	}

	// Output ISBN should be 13 digits
	expectedLength := 13
	actualLength := len(strings.TrimRight(outStream.String(), "\n"))
	if actualLength != expectedLength {
		t.Fatalf("Expected output length is %d but was %d.", expectedLength, actualLength)
	}
}

func TestRun_codeFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen --code 04", " ")

	status := cli.Run(args)

	// exitCode should be 0
	if status != exitCodeOK {
		t.Fatalf("Expected exit code is %d but was %d", exitCodeOK, status)
	}

	// Output ISBN should contain 9784(default prefix) + 04(pubcode)
	expected := "978404"
	if !strings.Contains(outStream.String(), expected) {
		t.Fatalf("Expected output contain %q but was %q", expected, outStream.String())
	}

	// Output ISBN should be 13 digits
	expectedLength := 13
	actualLength := len(strings.TrimRight(outStream.String(), "\n"))
	if actualLength != expectedLength {
		t.Fatalf("Expected output length is %d but was %d.", expectedLength, actualLength)
	}
}

func TestRun_codeFlagTooLong(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen --code 012345678", " ")

	status := cli.Run(args)

	// exitCode should be 0
	if status != exitCodeErr {
		t.Fatalf("Expected exit code is %d but was %d", exitCodeOK, status)
	}
}
