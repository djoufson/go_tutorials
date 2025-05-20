package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

const (
	Addition = iota
	Subtraction
	Multiplication
	Division
)

const (
	Red = "\033[31m"
	Reset = "\033[0m"
)

var operators = []string{"+", "-", "*", "/"}

func main() {
	fmt.Println("Welcome to my awesome calculator")
	var firstNumber = promptNumber("Insert the first number: ")
	var operator = promptOperator("Insert the operator: ")
	var secondNumber = promptNumber("Insert the second number: ")

	var result, err = calculate(firstNumber, secondNumber, operator)
	if err != nil {
		fmt.Printf("%v%v%v\n", Red, err, Reset)
		return
	}

	fmt.Print("\033[H\033[2J")
	fmt.Printf("%v %v %v = %v\n", firstNumber, operators[operator], secondNumber, result)
}

func calculate(firstNumber float64, secondNumber float64, operator int) (float64, error) {
	switch operator {
	case Addition:
		return firstNumber + secondNumber, nil
	case Subtraction:
		return firstNumber - secondNumber, nil
	case Multiplication:
		return firstNumber * secondNumber, nil
	case Division:
		if secondNumber == 0 {
			return 0, errors.New("can not divide by zero")
		}
		return firstNumber / secondNumber, nil
	default:
		return 0, errors.ErrUnsupported
	}
}

func validateOperators(input string) error {
	if slices.Contains(operators, input) {
		return nil
	}

	return errors.New("invalid operator")
}

func promptOperator(prompt string) int {
	var input string
	var err error = errors.ErrUnsupported
	for err != nil {
		fmt.Print(prompt)
		fmt.Scanln(&input)
		err = validateOperators(input)
		if err != nil {
			fmt.Printf("%vInsert a valid operator between + - * /%v\n", Red, Reset)
			continue
		}

		switch input {
		case "+":
			return Addition
		case "-":
			return Subtraction
		case "*":
			return Multiplication
		case "/":
			return Division
		default:
			err = errors.New("invalid operator selected")
			fmt.Printf("%vInsert a valid operator between + - * /%v\n", Red, Reset)
			continue
		}
	}

	return 0
}

func promptNumber(prompt string) float64 {
	var input string
	var err error = errors.ErrUnsupported
	for err != nil {
		fmt.Print(prompt)
		fmt.Scanln(&input)
		var value, conversionError = strconv.ParseFloat(input, 64)
		if conversionError != nil {
			fmt.Printf("%vInsert a valid number%v\n", Red, Reset)
			err = conversionError
			continue
		}

		return value
	}

	return 0
}
