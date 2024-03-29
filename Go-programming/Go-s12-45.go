package main

import (
	"fmt"
	// "math"
)

func main() {
	// var a [1024]byte
	// var s uint = 33

	// var i = 1 << s

	// fmt.Printf("i: %v.\n", i)

	// var j int32 = 1 << s

	// fmt.Printf("j: %v.\n", j)

	// fmt.Printf("3.1 * 3.1: %v.\n", 3.1 * 3.1)
	// fmt.Printf("3.16 * 3.16: %v.\n", 3.16 * 3.16)
	// fmt.Printf("3.162 * 3.162: %v.\n", 3.162 * 3.162)
	// fmt.Printf("3.1622 * 3.1622: %v.\n", 3.1622 * 3.1622)
	// fmt.Printf("3.16227 * 3.16227: %v.\n", 3.16227 * 3.16227)
	// fmt.Printf("3.162277 * 3.162277: %v.\n", 3.162277 * 3.162277)
	// fmt.Printf("1.41 * 1.41: %v.\n", 1.41 * 1.41)
	// fmt.Printf("1.414 * 1.414: %v.\n", 1.414 * 1.414)
	// fmt.Printf("1.4142 * 1.4142: %v.\n", 1.4142 * 1.4142)
	// fmt.Printf("1.41421 * 1.41421: %v.\n", 1.41421 * 1.41421)
	// fmt.Printf("1.414213 * 1.414213: %v.\n", 1.414213 * 1.414213)

	// fmt.Printf("math.Pi / 4: %v.\n", math.Pi / 4)
	// // fmt.Printf("0.25 * math.Pi: %v.\n", 0.25 * math.Pi)
	// fmt.Printf("math.Pi / 3: %v.\n", math.Pi / 3)
	// fmt.Printf("math.Pi / 2: %v.\n", math.Pi / 2)
	// // fmt.Printf("0.75 * math.Pi: %v.\n", 0.75 * math.Pi)
	// fmt.Printf("(3.0/4.0) * math.Pi: %v.\n\n", (3.0/4.0) * math.Pi)

	// fmt.Printf("math.Sin(0.5): %v.\n", math.Sin(0.5))
	// fmt.Printf("math.Sin(0.785): %v.\n", math.Sin(0.785))
	// fmt.Printf("math.Sin(1.0): %v.\n", math.Sin(1.0))
	// fmt.Printf("math.Sin(1.5): %v.\n", math.Sin(1.5))
	// fmt.Printf("math.Sin(2.0): %v.\n", math.Sin(2.0))
	// fmt.Printf("math.Sin(2.5): %v.\n", math.Sin(2.5))
	// fmt.Printf("math.Sin(3.0): %v.\n", math.Sin(3.0))

	// fmt.Printf("math.Exp(0.0) = %v.\n", math.Exp(0.0))
	// fmt.Printf("math.Exp(1.0) = %v.\n", math.Exp(1.0))
	// fmt.Printf("math.Exp(2.0) = %v.\n", math.Exp(2.0))
	// fmt.Printf("math.Exp(3.0) = %v.\n", math.Exp(3.0))
	// fmt.Printf("math.Exp(4.0) = %v.\n", math.Exp(4.0))
	// fmt.Printf("math.Exp(5.0) = %v.\n", math.Exp(5.0))
	// fmt.Printf("math.Exp(6.0) = %v.\n", math.Exp(6.0))
	// fmt.Printf("math.Exp(7.0) = %v.\n", math.Exp(7.0))
	// fmt.Printf("math.Exp(8.0) = %v.\n", math.Exp(8.0))
	// fmt.Printf("math.Exp(9.0) = %v.\n", math.Exp(9.0))
	// fmt.Printf("math.Exp(10.0) = %v.\n", math.Exp(10.0))

	// fmt.Printf("\n")

	// fmt.Printf("math.Pow(10, 0.3010299956) = %v.\n",
	// 	math.Pow(10, 0.3010299956))

	// fmt.Printf("T_eff = %v * 10^3 K.\n",
	// 	6.62607 * 2.99792 / (1.3806 * 4 * math.Pi))

	// fmt.Printf("300.0 / 1144.0 = %v.\n", 300.0 / 1144.0)

	// fmt.Printf("math.Pow(0.2622377, 4) / 3.0 = %v.\n",
	// 	math.Pow(0.2622377, 4) / 3.0)

	// fmt.Printf("math.Pow(0.2622, 4) / 3.0 = %v.\n",
	// 	math.Pow(0.2622, 4) / 3.0)

	// fmt.Printf("1.0/137: %v.\n", 1.0/137)
	// fmt.Printf("25 / 0.0254: %v.\n", 25 / 0.0254)

	// fmt.Printf("6350 * 0.0254: %v.\n", 6350 * 0.0254)

	// fmt.Printf("1 / math.E = %v.\n", 1 / math.E)

	// fmt.Printf("1 / (math.E * math.E) = %v.\n", 1 / (math.E * math.E))

	// fmt.Printf("7 * 33 = %v.\n", 7 * 33)

	// musicConst := math.Pow(2.0, 1.0/12.0)

	// fmt.Printf("musicConst: %v.\n", musicConst)

	// testVar := 1.0

	// for i := 1; i <= 12; i++ {
	// 	testVar *= 1.05

	// 	fmt.Printf("%v: %v.\n", i, testVar)
	// }

	approximation1 := 1.0594630945
	approximation2 := 1.0594630943

	result1 := 1.0
	result2 := 1.0

	for i := 1; i <= 12; i++ {
		result1 *= approximation1

		fmt.Printf("%v: %v.\n", i, result1)
	}

	difference := 2.0 - result1

	fmt.Printf("\n%v^12 = %v.\n\n", approximation1, result1)
	fmt.Printf("difference = %v.\n", difference)



	for i := 1; i <= 12; i++ {
		result2 *= approximation2

		fmt.Printf("%v: %v.\n", i, result2)
	}

	difference = 2.0 - result2

	fmt.Printf("\n%v^12 = %v.\n", approximation2, result2)
	fmt.Printf("difference = %v.\n", difference)

	// fmt.Printf("math.Pow(2.0, 1.0/12.0): %v.\n", math.Pow(2.0, 1.0/12.0))
}
