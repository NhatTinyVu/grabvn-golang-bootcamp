package main

import (
	"errors"
	"strings"
	"strconv"
)

func parser(expression string) (firstNumber float64, secondNumber float64, operator string, err error) {
	args := strings.Split(expression, " ")

	if (len(args) != 3) {
		err = errors.New("invalid expression format")
		return
	}

	firstNumber, err = strconv.ParseFloat(args[0], 64)
	if (err != nil) {
		err = errors.New("First number must be float")
		return
	}

	secondNumber, err = strconv.ParseFloat(args[2], 64)
	if (err != nil) {
		err = errors.New("Second number must be float")
		return
	}

	operator = args[1]

	return
}