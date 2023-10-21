package main

import (
	"fmt"
	// "golang.org/x/exp/constraints"
)

// type Ordered interface {
// 	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 |
// 		uint64 | uintptr | float32 | float64 | string
// }

// type Ordered interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
// 		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 |
// 		~float64 | ~string
// }

// type MyInt int

// func Max[E any](input []E) (max E) {
// 	for _, v := range input {
// 		if v > max {
// 			max = v
// 		}
// 	}

// 	return max
// }

// func Max[E Ordered](input []E) (max E) {
// 	for _, v := range input {
// 		if v > max {
// 			max = v
// 		}
// 	}

// 	return max
// }

// func Max[E constraints.Ordered](input []E) (max E) {
// 	for _, v := range input {
// 		if v > max {
// 			max = v
// 		}
// 	}

// 	return max
// }

// func main() {
// 	fmt.Printf("Max([]int{1, 2, 3}): %v.\n", Max([]int{1, 2, 3}))

// 	fmt.Printf("Max([]MyInt{MyInt(2), MyInt(4), MyIng(3)}): %v.\n",
// 		Max([]MyInt{MyInt(2), MyInt(4), MyInt(3)}))
// }

type Bunch[E any] []E

func PrintBunch[E any](b Bunch[E]) {
	for i, v := range b {
		fmt.Printf("Bunch %v: %v.\n", i, v)
	}
}

func (b Bunch[E]) Print() {
	fmt.Printf("Bunch: %v.\n", b)
}

type StingableBunch[E Stringer] []E

func main() {
	x := Bunch[int]{1, 2, 3}

	PrintBunch(x)
	x.Print()
}
