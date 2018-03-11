package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Fatalf("Expected exit code is %d but was %d", ExitCodeOK, status)
	}

	expected := fmt.Sprintf("isbn-gen version %s (%s)", Version, Commit)
	if !strings.Contains(errStream.String(), expected) {
		t.Fatalf("Expected output contain %q but was %q", expected, errStream.String())
	}
}

func TestRun_pubcodeFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen -p 04", " ")

	status := cli.Run(args)

	// ExitCode should be 0
	if status != ExitCodeOK {
		t.Fatalf("Expected exit code is %d but was %d", ExitCodeOK, status)
	}

	// Output ISBN should contain 9784(Japan code) + 04(pubcode)
	expected := fmt.Sprint("978404")
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

func TestRun_Repeat(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)

	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen -r 3", " ")

	status := cli.Run(args)

	if status != ExitCodeOK {
		t.Fatalf("Expected exit code is %d but was %d", ExitCodeOK, status)
	}

	expectedLength := 13*3 + 3 // 13 digits * 3 + 3 times \n
	actualLength := len(outStream.String())
	if actualLength != expectedLength {
		t.Errorf("Expected output length is %d but was %d.", expectedLength, actualLength)
	}
}

func TestRun_RepeatCannotBeZero(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)

	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen -r 0", " ")

	status := cli.Run(args)

	if status != ExitCodeError {
		t.Fatalf("Expected exit code is %d but was %d", ExitCodeError, status)
	}
}

func TestRun_RepeatCannotBeNegative(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)

	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen -r -1", " ")

	status := cli.Run(args)

	if status != ExitCodeError {
		t.Fatalf("Expected exit code is %d but was %d", ExitCodeError, status)
	}
}

func TestRun_RepeatOverMaximumValue(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)

	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./isbn-gen -r 100000001", " ")

	status := cli.Run(args)

	if status != ExitCodeError {
		t.Fatalf("Expected exit code is %d but was %d", ExitCodeError, status)
	}
}
