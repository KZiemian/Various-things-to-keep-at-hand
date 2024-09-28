package main

import (
	"fmt"
	"math"
)

func main() {
	x := 9.55

	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.6
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.65
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.7
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.75
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.8
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.85
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.9
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.95
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 10.0
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))
}
