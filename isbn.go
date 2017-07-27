package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Isbn struct {
	Number string
}

func NewIsbn(pubcode string) (*Isbn, error) {
	if len(pubcode) > 8 {
		return nil, fmt.Errorf("%s: pubcode must be less than 8 digits: %s", Name, pubcode)
	}
	return &Isbn{Number: generate(pubcode)}, nil
}

func (isbn *Isbn) GetNumber() string {
	return isbn.Number
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
