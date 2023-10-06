package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// func LastInt(s []int) int {
// 	return s[len(s) - 1]
// }

// func LastString(s []string) string {
// 	return s[len(s) - 1]
// }

// func main() {
// 	data1 := []int{1, 2, 3}
// 	fmt.Printf("LastInt(data1): %v.\n", LastInt(data1))

// 	data2 := []string{"a", "b", "c"}
// 	// fmt.Printf("LastInt(data2): %v.\n", LastInt(data2))
// 	fmt.Printf("LastString(data2): %v.\n", LastString(data2))
// }

// func Last[T any](s []T) T {
// 	return s[len(s) - 1]
// }

// func addComplex64(x, y complex64) complex64 {
// 	return x + y
// }


// func main() {
// 	data1 := []int{1, 2, 3}
// 	fmt.Printf("Last(data1): %v.\n", Last(data1))

// 	data2 := []string{"a", "b", "c"}
// 	fmt.Printf("Last(data2): %v.\n", Last(data2))
// }

func addComplex[T constraints.Complex](x, y T) T {
	return x + y
}

func main() {
	// fmt.Printf("addComplex64(complex64(1), complex64(2)): %v.\n",
	// 	addComplex64(complex64(1), complex64(2)))

	// fmt.Printf("addComplex64(complex128(1), complex128(2)): %v.\n",
	// 	addComplex64(complex128(1), complex128(2)))

	fmt.Printf("addComplex(complex64(1), complex64(2)): %v.\n",
		addComplex(complex64(1), complex64(2)))

	fmt.Printf("addComplex(complex128(1), complex128(2)): %v.\n",
		addComplex(complex128(1), complex128(2)))
}

// type Complex interface {
// 	~complex64 | ~complex128
// }

// type Float interface {
// 	~float32 | ~float64
// }

// type Integer interface {
// 	Signed | Unsigned
// }

// type Ordered interface {
// 	Integer | Float | ~string
// }

// type Signed interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64
// }

// type Unsigned interface {
// 	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
// }

type Doer interface {
	DoSomething()
}

func Last[T Doer](s []T) T {
	return s[len(s) - 1]
}

type Integer interface {
	int
}

type Number interface {
	int | float64
}

type Number interface {
	constraints.Integer | constraints.Float
}
