package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}

	return y
}

func main() {
	fmin := min[float64]

	m := fmin(2.71, 3.14)

	fmt.Printf("m: %v.\n", m)
}
