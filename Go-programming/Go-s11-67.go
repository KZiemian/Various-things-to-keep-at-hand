package main

import "fmt"

func main() {
	a := 7
	b := 2

	fmt.Println("a & b:", a & b)
	fmt.Println("a | b:", a | b)
	fmt.Println("a ^ b:", a ^ b)
	fmt.Println("a &^ b:", a &^ b)
	fmt.Println("a << 1:", a << 1)
	fmt.Println("a << 2:", a << 2)
	fmt.Println("a >> 1:", a >> 1)
}
