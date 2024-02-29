package main

import "fmt"

// func main() {
// 	factorialInt := 1
// 	for i := 1; i <= 21; i++ {
// 		fmt.Printf("%v! = %v.\n", i-1, factorialInt)

// 		factorialInt *= i
// 	}
// }
// 20! = 2_432_902_008_176_640_000.

func main() {
	eApprox := 0.0
	n := 7
	term := 1.0

	eApprox += term

	for i := 1; i <= n; i++ {
		term /= float64(i)
		eApprox += term
	}

	fmt.Printf("Po wysumowaniu %v wyrazów, eApprox = %v.\n", n + 1,
		eApprox)
}
