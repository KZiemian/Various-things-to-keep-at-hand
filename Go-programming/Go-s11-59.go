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
	fmt.Printf("min(1, 2): %v.\n", min(1, 2))
	fmt.Printf("min(3, 2): %v.\n", min(3, 2))
	fmt.Printf("min(0.0, 1.0): %v.\n", min(0.0, 1.0))
	fmt.Printf("min(5.0, 3.0): %v.\n", min(5.0, 3.0))
}
