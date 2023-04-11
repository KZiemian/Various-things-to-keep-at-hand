package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Printf("sum(1, 2): %v\n", sum(1, 2))
}
