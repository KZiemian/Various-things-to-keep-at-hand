package main

import (
	"fmt"
)

type weekday int

type ordered interface {
	int | float64 | weekday
}

const (
	Monday weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(max([]weekday{Monday, Tuesday, Sunday}))
	fmt.Println(max([]weekday{Sunday, Friday}))
	fmt.Println(max([]weekday{Friday, Monday}))
}

func max[T ordered](input []T) T {
	var max T

	for _, v := range input {
		if v > max {
			max = v
		}
	}

	return max
}
