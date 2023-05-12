package main

import (
	"fmt"
	"math"
	"gonum.org/v1/gonum/integrate"
)

func main() {
	// valuesOfFun := make([]float64, 9)
	// dx := 1.0/8.0
	// x := 0.0

	// for i := 0; i < 9; i++ {
	// 	valuesOfFun[i] = x * x

	// 	x += dx
	// }

	// // fmt.Printf("valuesOfFun: %v.\n", valuesOfFun)
	// integralValue := integrate.Romberg(valuesOfFun, dx)

	// fmt.Printf("integralValue: %v.\n", integralValue)

	valuesOfFun := make([]float64, 129)
	dx := 3.1415/(2.0 * 129)
	x := 0.0

	for i := 0; i < 129; i++ {
		valuesOfFun[i] = math.Sin(x)

		x += dx
	}
	// fmt.Printf("valuesOfFun: %v.\n", valuesOfFun)

	integralValue := integrate.Romberg(valuesOfFun, dx)

	fmt.Printf("integralValue: %v.\n", integralValue)
}
