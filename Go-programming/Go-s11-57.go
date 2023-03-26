package main

import "fmt"

func main() {
	powerOfTwo := 1

	fmt.Printf("2^0 = 1.\n")
	
	for i := 1; i <= 20; i++ {
		powerOfTwo *= 2
		fmt.Printf("2^%v = %v.\n", i, powerOfTwo)
	}
}
