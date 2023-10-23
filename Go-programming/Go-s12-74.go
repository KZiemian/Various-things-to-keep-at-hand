package main

// import (
// 	"fmt"
// 	// "golang.org/x/exp/constraints"
// )

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

// type Bunch[E any] []E

// func PrintBunch[E any](b Bunch[E]) {
// 	for i, v := range b {
// 		fmt.Printf("Bunch %v: %v.\n", i, v)
// 	}
// }

// func (b Bunch[E]) Print() {
// 	fmt.Printf("Bunch: %v.\n", b)
// }

// type StingableBunch[E Stringer] []E

// func main() {
// 	x := Bunch[int]{1, 2, 3}

// 	PrintBunch(x)
// 	x.Print()
// }

// func TypeOf(i interface{}) Type

func (v Value) Interface() interface{}

import (
	"fmt"
	"reflect"
)

func main() {
	// var x float64 = 3.4

	// fmt.Println("value:", reflect.ValueOf(x))
	// fmt.Println("value:", reflect.ValueOf(x).String())

	// v := reflect.ValueOf(x)

	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value:", v.Float())

	var x uint8 = 'x'

	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	fmt.Println("kind is uint8:", v.Kind() == reflect.Uint8)
	x = uint8(v.Uint())

	// y := v.Interface().(float64)

	fmt.Println(y)
}
