package main

import (
	"fmt"
)

type weekday int

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

	type ordered interface {
		int | float64 | weekday
	}
}

type ordered interface {
	int | float64 | weekday
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
