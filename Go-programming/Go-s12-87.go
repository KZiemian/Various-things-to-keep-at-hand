package main

import (
	"fmt"
	"math"
)

type Calculator struct {
	acc float64
}

type opfunc func(float64, float64) float64

// func (c *Calculator) Do(op opfunc, v float64) float64 {
// 	c.acc = op(c.acc, v)

// 	return c.acc
// }

// func Add(a, b float64) float64 {
// 	return a + b
// }

// func Sub(a, b float64) float64 {
// 	return a - b
// }

// func Mul(a, b float64) float64 {
// 	return a * b
// }

// func main() {
// 	var c Calculator

// 	fmt.Printf("c.Do(Add, 5): %v.\n", c.Do(Add, 5))
// 	fmt.Printf("c.Do(Sub, 3): %v.\n", c.Do(Sub, 3))
// 	fmt.Printf("c.Do(Mul, 8): %v.\n", c.Do(Mul, 8))
// }

// func Sqrt(n, _ float64) float64 {
// 	return math.Sqrt(n)
// }

// func main() {
// 	var c Calculator

// 	fmt.Printf("c.Do(Add, 16): %v.\n", c.Do(Add, 16))
// 	fmt.Printf("c.Do(Sqrt, 0): %v.\n", c.Do(Sqrt, 0))
// }

func Add(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc + n
	}
}

func Sub(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc - n
	}
}

func Mul(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc * n
	}
}

func (c *Calculator) Do(op func(float64) float64) float64 {
	c.acc = op(c.acc)

	return c.acc
}

func Sqrt() func(float64) float64 {
	return func(n float64) float64 {
		return math.Sqrt(n)
	}
}

func main() {
	var c Calculator

	// fmt.Printf("c.Do(Add(10)): %v.\n", c.Do(Add(10)))
	// fmt.Printf("c.Do(Add(20)): %v.\n", c.Do(Add(20)))
	// fmt.Printf("c.Do(Sub(5)):  %v.\n", c.Do(Sub(5)))
	// fmt.Printf("c.Do(Mul(3)):  %v.\n", c.Do(Mul(3)))
	// fmt.Printf("c.Do(Sqrt()):  %v.\n", c.Do(Sqrt()))

	// fmt.Printf("c.Do(Add(2)): %v.\n", c.Do(Add(2)))
	// fmt.Printf("c.Do(Sqrt()): %v.\n", c.Do(Sqrt()))

	fmt.Printf("c.Do(Add(2)):    %v.\n", c.Do(Add(2)))
	fmt.Printf("c.Do(math.Sqrt): %v.\n", c.Do(math.Sqrt))
	fmt.Printf("c.Do(math.Cos):  %v.\n", c.Do(math.Cos))
}

type Mux struct {
	mu    sync.Mutex
	conns map[net.Addr]net.Conn
}
