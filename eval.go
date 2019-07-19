package main

import (
	"errors"
)

func eval(expression string) (result float64, err error) {
	firstNumber, secondNumber, operator, err := parser(expression)

	if (err != nil) {
		return
	}

	switch operator {
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber - secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			if secondNumber == 0 {
				err = errors.New("cannot divide by zero")
			} else {
				result = firstNumber / secondNumber
			}
		default:
			err = errors.New("Not supported operator")
	}

	return
}