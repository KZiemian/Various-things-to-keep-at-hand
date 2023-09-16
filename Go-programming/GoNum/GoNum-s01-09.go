package main

import (
	"fmt"
	"math/rand"
	"time"
	// "gonum.org/v1/gonum/dist"
	"gonum.org/v1/gonum/distuv"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	dist := distuv.ChiSquared{K: 10, Src: rnd}
	nSamp := 1000

	x := make([]float64, nSamp)
	for i := range x {
		x[i] = dist.Rand()
	}

	mean := stat.Mean(x, nil)

	fmt.Println("True mean:", dist.Mean())
	fmt.Println("Estimated mean:", mean)
}
