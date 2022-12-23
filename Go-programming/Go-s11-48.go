package main

import (
	"fmt"
)

func main() {
	fmt.Println(max([]int{1, 2, 3}))
	fmt.Println(max([]int{3, 2, 1}))
}

type ordered interface {
	int | float64
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
