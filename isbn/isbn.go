package isbn

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ISBN represents ISBN code.
type ISBN struct {
	Number string
}

// NewISBN generate ISBN struct with valid ISBN code.
func NewISBN(group, pubcode string) (*ISBN, error) {
	identifier := SearchIdentifier(group)
	if identifier == nil {
		return nil, fmt.Errorf("%q is not supported.", group)
	}
	if !isNumber(pubcode) {
		return nil, fmt.Errorf("pubcode must be a number: %s", pubcode)
	}
	if len(identifier.Prefix+pubcode) > 12 {
		return nil, fmt.Errorf("prefix + pubcode must be equal or less than 12 digits: prefix=%s, pubCode=%s",
			identifier.Prefix, pubcode)
	}
	return &ISBN{Number: generate(identifier.Prefix, pubcode)}, nil
}

// Generates 13 digits which is valid as ISBN code.
func generate(prefix, pubcode string) string {
	rand.Seed(time.Now().UnixNano())

	isbn := prefix + pubcode
	rest := 12 - len(isbn)
	for i := 0; i < rest; i++ {
		isbn += strconv.Itoa(rand.Intn(10))
	}
	return isbn + calcCheckDigit(isbn)
}

// Calculate ISBN last digit which is called check digit.
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

func isNumber(pubcode string) bool {
	if len(pubcode) == 0 {
		return true
	}
	if _, err := strconv.Atoi(pubcode); err == nil {
		return true
	}
	return false
}
