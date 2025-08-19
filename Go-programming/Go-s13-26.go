package main

import (
	"fmt"
	"math"
)

func main() {
	// for x := 19.5; x <= 20.0; x += 0.01 {
	// 	fmt.Printf("math.J0(%v) = %v.\n", x, math.J0(x))
	// }


	// x := 20.0
	// fmt.Printf("math.J0(%v) = %v.\n", x, math.J0(x))

	// for x := 3.4; x <= 3.5; x += 0.01 {
	// 	fmt.Printf("math.Log2(%v) = %v.\n", x, math.Log2(x))
	// }

	// x := 0.001
	// fmt.Printf("math.Log2(%v) = %v.\n", x, math.Log2(x))

	// x := 0.001
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// for x := 19.8; x <= 20.01; x += 0.01 {
	// 	fmt.Printf("math.J1(%v) = %v.\n", x, math.J1(x))
	// }

	for x := 19.4; x <= 20.0; x += 0.01 {
		fmt.Printf("math.Jn(2, %v) = %v.\n", x, math.Jn(2, x))
	}

	// x := 0.001
	// fmt.Printf("math.Log10(%v) = %v.\n", x, math.Log10(x))

	// x := 20.0
	// fmt.Printf("math.Jn(2, %v) = %v.\n", x, math.Jn(2, x))
}
