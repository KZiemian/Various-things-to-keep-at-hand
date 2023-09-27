package main

import (
	"fmt"
	// "math"
)

// func main() {
// 	musicConst := 1.0594630943592953

// 	fmt.Printf("musicConst: %v.\n", musicConst)

// 	difference := musicConst - math.Pow(2.0, 1.0/12.0)

// 	fmt.Printf("difference: %v.\n", difference)

// 	testVar := 1.0

// 	for i := 1; i <= 12; i++ {
// 		testVar *= musicConst

// 		fmt.Printf("%v: %v.\n", i, testVar)
// 	}
// }

func main() {
	testedValue := 1.0594630943

	powerOfTestedValue := 1.0

	for i := 1; i <= 12; i++ {
		powerOfTestedValue *= testedValue

		fmt.Printf("%v: %v.\n", i, powerOfTestedValue)
	}
}
