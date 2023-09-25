package main

import (
	"fmt"
	"errors"
)

func main() {
	operator := "+"
	nr1 := 10.0
	nr2 := 2.0

	result, err := calculate(operator, nr1, nr2)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())

		return
	}

	fmt.Printf("The result is: %f.\n\n", result)


	operator = "-"

	result, err = calculate(operator, nr1, nr2)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())

		return
	}

	fmt.Printf("The result is: %f.\n\n", result)


	operator = "*"

	result, err = calculate(operator, nr1, nr2)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())

		return
	}

	fmt.Printf("The result is: %f.\n\n", result)


	nr2 = 0.0

	operator = "/"

	result, err = calculate(operator, nr1, nr2)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())

		return
	}

	fmt.Printf("The result is: %f.\n", result)
}



func calculate(operator string, nr1, nr2 float64) (float64, error) {
	if operator == "+" {
		return add(nr1, nr2), nil
	} else if operator == "-" {
		return subtract(nr1, nr2), nil
	} else if operator == "*" {
		return multiply(nr1, nr2), nil
	} else if operator == "/" {
		return divide(nr1, nr2)
	} else {
		return 0.0, fmt.Errorf("no calculation for operator %s",
			operator)
	}
}

func add(nr1, nr2 float64) float64 {
	return nr1 + nr2
}

func subtract(nr1, nr2 float64) float64 {
	return nr1 - nr2
}

func multiply(nr1, nr2 float64) float64 {
	return nr1 * nr2
}

func divide(nr1, nr2 float64) (float64, error) {
	if nr2 == 0.0 {
		return 0.0, errors.New("division by zero")
	}

	return nr1 / nr2, nil
}
