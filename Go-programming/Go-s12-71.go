package main

import (
	"fmt"
	"math"
)

// func Add(a int, b int) int {
// 	return a + b
// }

// func AddInt(a int, b int) int {
// 	return a + b
// }

// func AddFloat64(a float64, b float64) float64 {
// 	return a + b
// }

// type Num interface {
// 	int | int8 | int16 | float32 | float64
// }

// type UserID int

// func Add[T ~int | float64](a T, b T) T {
// 	return a + b
// }

// func main() {
	// result := Add(1, 2)

	// fmt.Printf("Add(1, 2): %v.\n", result)

	// result = Add(a: 2, b: 2)

	// fmt.Printf("Add(a: 2, b: 2): %v.\n", result)

	// result = Add(1.1, 2.2)
	// fmt.Printf("Add(1.1, 2.2): %v.\n", result)

	// result := AddInt(1, 2)

	// fmt.Printf("AddInt(1, 2): %v.\n", result)

	// result = AddInt(1.1, 2.2)

	// fmt.Printf("AddInt(1.1, 2.2): %v.\n", result)

	// anotherResult := AddFloat64(1.1, 2.2)

	// fmt.Printf("AddFloat64(1.1, 2.2): %v.\n", anotherResult)

	// result1 := Add(1, 2)

	// fmt.Printf("result1: %v, %T.\n", result1, result1)

	// result2 := Add(1.1, 2.2)

	// fmt.Printf("result2: %v, %T.\n", result2, result2)

	// result1 := Add(1, 2)

	// fmt.Printf("result1: %T, %v.\n", result1, result1)

	// result2 := Add(1.1, 2.2)

	// fmt.Printf("result2: %T, %v.\n", result2, result2)

	// a := UserID(1)
	// b := UserID(2)

	// result := Add(a, b)

	// fmt.Printf("result: %v.\n", result)
// }

func MapValues(values []int, mapFunc func(int) int) []int {
	var newValues []int

	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}

	return newValues
}

func main() {
	result := MapValues([]int{1, 2, 3}, func(n int) int {
		return n * 2
	})

	fmt.Printf("result: %v.\n", result)

	result = MapValues([]int{3, 5, 7}, func(n int) int {
		return n * 3
	})

	fmt.Printf("result: %v.\n", result)

	resultAnother := MapValues([]int{0, 1, 2, 3}, math.Sin)

	fmt.Printf("result: %v.\n", resultAnother)
}
