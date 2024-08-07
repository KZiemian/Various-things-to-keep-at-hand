package main

import (
	"fmt"
	"math"
)

// func main() {
// 	factorialInt := 1
// 	for i := 1; i <= 21; i++ {
// 		fmt.Printf("%v! = %v.\n", i-1, factorialInt)

// 		factorialInt *= i
// 	}
// }
// 20! = 2_432_902_008_176_640_000.

// func main() {
// 	eApprox := 0.0
// 	n := 7
// 	term := 1.0

// 	eApprox += term

// 	for i := 1; i <= n; i++ {
// 		term /= float64(i)
// 		eApprox += term
// 	}

// 	fmt.Printf("Po wysumowaniu %v wyrazów, eApprox = %v.\n", n + 1,
// 		eApprox)
// }

func main() {
	// alpha := 0.00729735
	// alphaInv := 137.035

	// fmt.Printf("1 / alpha = %v.\n", 1 / alpha)
	// fmt.Printf("1 / alphaInv = %v.\n", 1 / alphaInv)

	// fmt.Printf("1.0 / 137.0 = %v.\n", 1.0 / 137.0)

	// fmt.Printf("4 * math.Pi = %v.\n", 4 * math.Pi)

	// liczbaWierszyPanaTadeusza := 985 + 850 + 789 + 1002 + 905 + 616
	// liczbaWierszyPanaTadeusza += 554 + 805 + 762 + 901 + 681 + 863

	// fmt.Printf("Liczba wierszy „Pana Tadeusza”: %v.\n",
	// 	liczbaWierszyPanaTadeusza)

	// fmt.Printf("math.Sqrt2/2.0 = %v.\n", math.Sqrt2/2.0)
	// fmt.Printf("math.Sqrt2/3.0 = %v.\n", math.Sqrt2/3.0)
	// fmt.Printf("math.Sqrt2/4.0 = %v.\n", math.Sqrt2/4.0)

	sqrt3 := math.Sqrt(3)

	fmt.Printf("sqrt3/2.0 = %v.\n", sqrt3/2.0)
	fmt.Printf("sqrt3/3.0 = %v.\n", sqrt3/3.0)
	fmt.Printf("sqrt3/4.0 = %v.\n", sqrt3/4.0)
}
