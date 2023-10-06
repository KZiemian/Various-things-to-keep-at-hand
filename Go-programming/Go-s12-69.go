package main

import (
	"fmt"
	"sort"
)

func median(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)

	sort.Float64s(dataCopy)

	var median float64
	l := len(dataCopy)

	if l == 0 {
		return 0
	} else if l % 2 == 0 {
		median = (dataCopy[l/2 - 1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	return median
}

func main() {
	fmt.Println(median([]float64{1, 3, 2}))
	fmt.Println(median([]float64{1}))
	fmt.Println(median([]float64{3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1,
		1, 1, 111}))
}
