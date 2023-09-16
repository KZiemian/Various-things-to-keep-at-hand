package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main() {
	x := []float64{7, 5, 3, 1, 2}

	// Computer raw mean and standard deviation.
	mean, std := stat.MeanStdDev(x, nil)

	fmt.Printf("unweighted mean = %v, std = %v.\n", mean, std)

	// Compute weighted mean and standard deviation.
	weights := []float64{2, 6, 1, 9, 2}
	mean, std = stat.MeanStdDev(x, weights)
	fmt.Printf("weighted mean = %v, std = %v\n", mean, std)

	// Compute weighted median.
	stat.SortWeighted(x, weights)

	fmt.Println("sorted x =", x)
	fmt.Println("sorted w =", weights)

	median := stat.Quantile(0.5, stat.Empirical, x, weights)
	fmt.Println("median =", median)
}
