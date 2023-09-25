package main

import (
	"fmt"
	// "math/rand"
	"gonum.org/v1/gonum/mat"
)

// func main() {
// 	zero := mat.NewDense(3, 5, nil)

// 	fmt.Printf("zero: %v.\n", zero)
// 	fmt.Printf("zero: %T.\n", zero)
// }

// func main() {
// 	data := make([]float64, 36)

// 	for i := range data {
// 		data[i] = rand.NormFloat64()
// 	}

// 	a := mat.NewDense(6, 6, data)

// 	fmt.Printf("a: %v.\n", a)
// 	fmt.Printf("a: %T.\n", a)
// }

// func main() {
// 	data := []float64{1, 0, 3,
// 		0,  2, 0,
// 		0, -1, 3}

// 	a := mat.NewDense(3, 3, data)

// 	fmt.Printf("a: %v.\n", a)
// 	fmt.Printf("a: %T.\n", a)

// 	tr := mat.Trace(a)

// 	fmt.Printf("tr: %v.\n", tr)
// }

func main() {
	data := []float64{1, 0, 3,
		0,  2, 0,
		0, -1, 3}

	a := mat.NewDense(3, 3, data)

	var c mat.Dense
	c.Mul(a, a)

	fmt.Printf("c: %v.\n", c)
}
