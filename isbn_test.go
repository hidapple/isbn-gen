package main

import (
	"testing"
)

func TestNewISBN_Init(t *testing.T) {
	tests := []struct {
		groupCode string
		pubCode   string
	}{
		{"jp", ""},
		{"jp", "00000000"},
		{"br1", "0000000"},
	}

	for i, tt := range tests {
		if _, err := NewISBN(tt.groupCode, tt.pubCode); err != nil {
			t.Errorf("Case[%d]: ISBN is expected to be able to init with grupCode=%q, pubCode=%q",
				i, tt.groupCode, tt.pubCode)
		}
	}
}

func TestNewISBN_CannotInit(t *testing.T) {
	tests := []struct {
		groupCode string
		pubCode   string
	}{
		{"jp", "000000000"},
		{"br1", "00000000"},
	}

	for i, tt := range tests {
		if _, err := NewISBN(tt.groupCode, tt.pubCode); err == nil {
			t.Errorf("Case[%d]: ISBN should not be able to init with groupCode=%q, pubCode=%q",
				i, tt.groupCode, tt.pubCode)
		}
	}
}
