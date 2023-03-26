package main

import "fmt"

func main() {
	b1 := true
	b2 := false

	fmt.Printf("b1: %v.\n", b1)
	fmt.Printf("!b1: %v.\n", !b1)
	fmt.Printf("b2: %v.\n", b2)
	fmt.Printf("!b2: %v.\n", !b2)
	fmt.Printf("b1 == b2: %v.\n", b1 == b2)
	fmt.Printf("b1 != b2: %v.\n", b1 != b2)
	fmt.Printf("b1 && b2: %v.\n", b1 && b2)
	fmt.Printf("b1 || b2: %v.\n", b1 || b2)
}
