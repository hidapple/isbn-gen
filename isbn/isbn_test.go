package isbn

import (
	"testing"
)

func TestNewISBN_Init(t *testing.T) {
	tests := []struct {
		groupCode string
		code      string
	}{
		{"Japan", ""},
		{"jp", ""},
		{"jp", "00000000"},
		{"br1", "0000000"},
	}

	for i, tt := range tests {
		if _, err := NewISBN(tt.groupCode, tt.code); err != nil {
			t.Errorf("case %d: ISBN is expected to be able to init with grupCode=%q, code=%q",
				i, tt.groupCode, tt.code)
		}
	}
}

func TestNewISBN_CannotInit(t *testing.T) {
	tests := []struct {
		groupCode string
		code      string
	}{
		{"Japan", "000000000"},
		{"jp", "000000000"},
		{"br1", "00000000"},
	}

	for i, tt := range tests {
		if _, err := NewISBN(tt.groupCode, tt.code); err == nil {
			t.Errorf("case %d: ISBN should not be able to init with groupCode=%q, code=%q",
				i, tt.groupCode, tt.code)
		}
	}
}

func TestNumber(t *testing.T) {
	isbn, err := NewISBN("jp", "")
	if err != nil {
		t.Fatalf("failed initializing ISBN")
	}

	actual := isbn.Number()
	if len(actual) != 13 {
		t.Fatalf("expect Number() returns 13 digit numbers but got %q", actual)
	}
}
