package main

import "fmt"

func main() {
	a := 7
	b := 2

	// a = b
	// fmt.Println("a:", a)

	// a = 7
	// a += b
	// fmt.Println("a:", a)

	// a = 7
	// a -= b
	// fmt.Println("a:", a)

	// a = 7
	// a *= b
	// fmt.Println("a:", a)

	// a = 7
	// a /= b
	// fmt.Println("a:", a)

	// a = 7
	// a %= b
	// fmt.Println("a:", a)

	// a = 7
	// a &= b
	// fmt.Println("a:", a)

	// a = 7
	// a |= b
	// fmt.Println("a:", a)

	// a = 7
	// a ^= b
	// fmt.Println("a:", a)

	result := false && a*b + a - 21 == 32 || true
	fmt.Println(result)
}
