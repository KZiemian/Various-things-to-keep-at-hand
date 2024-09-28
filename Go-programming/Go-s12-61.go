// package main

// import (
// 	"fmt"
// 	"math"
// )

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

// package main

// import (
// 	"fmt"

// 	"golang.org/x/exp/constraints"
// 	"golang.org/x/exp/slices"
// )

// type Number interface {
// 	constraints.Float | constraints.Integer
// }

// func median[T Number](data []T) float64 {
// 	dataCopy := make([]T, len(data))
// 	copy(dataCopy, data)

// 	slices.Sort(dataCopy)

// 	var median float64
// 	l := len(dataCopy)

// 	if l == 0 {
// 		return 0
// 	} else if l % 2 == 0 {
// 		median = float64((dataCopy[l/2 - 1] + dataCopy[l/2]) / 2.0)
// 	} else {
// 		median = float64(dataCopy[l/2])
// 	}

// 	return median
// }

// func main() {
// 	fmt.Println(median([]float64{1, 3, 2}))
// 	fmt.Println(median([]int{1}))
// 	fmt.Println(median([]int32{3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1,
// 		1, 1, 111}))
// }

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.55

	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.6
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.65
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.7
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.75
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.8
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.85
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.9
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.95
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.0
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))



	// x = 1.05
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.1
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.15
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.2
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.25
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.3
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.35
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.4
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.45
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 1.5
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 1.55
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 1.6
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 1.65
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 1.7
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 1.75
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 1.8
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 1.85
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 1.9
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 1.95
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.0
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))
}
