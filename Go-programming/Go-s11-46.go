package main

import (
	"fmt"
)

func sumInts(v1, v2 int) int {
	return v1 + v2
}

func sumFloat64s(v1, v2 float64) float64 {
	return v1 + v2
}

func sumStrings(v1, v2 string) string {
	return v1 + v2
}

func main() {
	fmt.Printf("sumInts(1, 2): %v\n", sumInts(1, 2))
	fmt.Printf("sumFloat64s(1.0, 2.0): %v\n", sumFloat64s(1.0, 2.0))
	fmt.Printf("sumStrings(\"1\", \"2\"): %v\n", sumStrings("1", "2"))
}
