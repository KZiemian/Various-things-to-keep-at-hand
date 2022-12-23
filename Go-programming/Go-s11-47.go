package main

import (
	"constraints"
	"fmt"
)

func sumFunc[T constraints.Integer](v1, v2 T) T {
	return v1 + v2
}

func main() {
	fmt.Printf("sumFunc(1, 2): %v\n", sumFunc(1, 2))
	fmt.Printf("sumFunc(1.0, 2.0): %v\n", sumFunc(1.0, 2.0))
	fmt.Printf("sumFunc(\"1\", \"2\")", sumFunc("1", "2"))
}
