package main

import (
	"testing"
)

func TestNewIsbn_Init(t *testing.T) {
	// when pubcode is not specified
	p1 := ""
	_, err1 := NewIsbn(p1)
	if err1 != nil {
		t.Error("Isbn is expected to be able to init with empty pubcode")
	}

	// when pubcode is less than 8 digits
	p2 := "00000000"
	_, err2 := NewIsbn(p2)
	if err2 != nil {
		t.Errorf("Isbn is expected to be able to init with %d digits pubcode", len(p2))
	}
}

func TestNewIsbn_CannotInit(t *testing.T) {
	// when pubcode is 8 digits or more
	p1 := "000000000"
	_, err1 := NewIsbn(p1)
	if err1 == nil {
		t.Errorf("Isbn should not be able to init with %d digits pubcode", len(p1))
	}

	// when pubcode is not number
	p2 := "abc"
	_, err2 := NewIsbn(p1)
	if err2 == nil {
		t.Errorf("Isbn should not be able to init with not a number pubcode: %s", p2)
	}
}

func TestGetNumber(t *testing.T) {
	expected := "9784060370112"
	isbn := &Isbn{Number: expected}

	actual := isbn.GetNumber()
	if actual != expected {
		t.Errorf("Expected is %s, but was %s", expected, actual)
	}
}
