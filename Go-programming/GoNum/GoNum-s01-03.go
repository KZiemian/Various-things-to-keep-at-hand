package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	d := mat.NewDense(3, 4, nil)
	fmt.Printf("%v\n\n", mat.Formatted(d))

	data := []float64{
		6, 3, 5,
		-1, 9, 7,
		2, 3, 4,
	}

	d2 := mat.NewDense(3, 3, data)
	fmt.Printf("%v\n", mat.Formatted(d2))
}
