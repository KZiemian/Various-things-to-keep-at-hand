package main

import (
	"fmt"
	"math"
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

// func main() {
// 	testedValue := 1.0594630943

// 	powerOfTestedValue := 1.0

// 	for i := 1; i <= 12; i++ {
// 		powerOfTestedValue *= testedValue

// 		fmt.Printf("%v: %v.\n", i, powerOfTestedValue)
// 	}
// }

// func constains[T comparable](elems []T, v T) bool {
// 	for _, s := range elems {
// 		if v == s {
// 			return true
// 		}
// 	}

// 	return false
// }

// func Last[T constraints.Complex](s []T) {}
// func Last[T constraints.Float](s []T) {}
// func Last[T constraints.Integer](s []T) {}
// func Last[T constraints.Ordered](s []T) {}
// func Last[T constraints.Signed](s []T) {}
// func Last[T constraints.Unsigned](s []T) {}

func main() {
	value := 0.25*(1 + 1 / (math.E * math.E)) - 5.0 / (4.0 * math.E * math.E)

	fmt.Printf("value: %v.\n", value)
}
