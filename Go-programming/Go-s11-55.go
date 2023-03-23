package main

import (
	"fmt"
)

func main() {
	// s := "hello"
	// c := []byte(s)

	// fmt.Printf("s, %v: %T.\n", s, s)
	// fmt.Printf("c, %v: %T.\n", c, c)

	// var strVar string

	// strVar = string(c)

	// fmt.Printf("strVar, %v: %T.\n", strVar, strVar)

	// i := 0

	// i += 15
	// i += 15 * 16
	// i += 15 * 16 * 16
	// i += 15 * 16 * 16 * 16

	// fmt.Printf("Wartosc 0xFFFF: %v.\n", i)

	i := 0

	i += 15
	i += 15 * 16
	i += 14 * 16 * 16
	i += 15 * 16 * 16 * 16

	fmt.Printf("Wartosc 0xFEFF: %v.\n", i)
}
