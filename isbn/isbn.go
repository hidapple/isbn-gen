package isbn

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ISBN represents ISBN number.
type ISBN struct {
	// prefix is prefix of ISBN.
	prefix string

	// identifier is identifying group of ISBN.
	identifier string

	// code is code part of ISBN which consist of publisher code and book code.
	code string

	// checkDigit is check digit value of ISBN.
	checkDigit string
}

// NewISBN generate ISBN struct with valid ISBN code.
func NewISBN(group, code string) (*ISBN, error) {
	id := SearchIdentifier(group)
	if id == nil {
		return nil, fmt.Errorf("%q is not supported", group)
	}
	if !isNumber(code) {
		return nil, fmt.Errorf("bookCode must be a number: %s", code)
	}
	if len(id.Prefix+id.Identifier+code) > 12 {
		return nil, fmt.Errorf("code is too long: code=%s", code)
	}

	code = generateCode(code, 12-len(id.Prefix+id.Identifier))
	checkDigit := calcCheckDigit(id.Prefix + id.Identifier + code)
	return &ISBN{
		prefix:     id.Prefix,
		identifier: id.Identifier,
		code:       code,
		checkDigit: checkDigit,
	}, nil
}

// Number returns full ISBN number.
func (isbn *ISBN) Number() string {
	return isbn.prefix + isbn.identifier + isbn.code + isbn.checkDigit
}

// generateCode generates ISBN code part with given prefix.
func generateCode(prefix string, length int) string {
	rand.Seed(time.Now().UnixNano())

	code := prefix
	for i := 0; i < length-len(prefix); i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}

// calcCheckDigit calculates ISBN last digit which is called check digit.
func calcCheckDigit(isbn string) string {
	sum := 0
	for i, v := range strings.Split(isbn, "") {
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
	}
	return strconv.Itoa(calcResult)
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
