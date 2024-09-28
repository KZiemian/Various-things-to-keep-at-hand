// package main

// import "fmt"

// // func Last[T interface{ int | int8 | int16 | int32 }](s []T) T {
// // 	return s[len(s) - 1]
// // }

// func Last[T int | int8 | int16 | int32 ](s []T) T {
// 	return s[len(s) - 1]
// }

// func main() {
// 	data1 := []int{1, 2, 3}
// 	data2 := []int8{-1, -2, -3}

// 	fmt.Printf("Last(data1): %v.\n", Last(data1))
// 	fmt.Printf("Last(data2): %v.\n", Last(data2))
// }

// package main

// import "fmt"

// type MyInt int

// type Int interface {
// 	~int | int8 | int16 | int32
// }

// func Last[T Int](s []T) T {
// 	return s[len(s) - 1]
// }

// func main() {
// 	data := []MyInt{1, 2, 3}

// 	fmt.Println(Last(data))
// }

// package main

// import "fmt"

// type KV[K comparable, V any] struct {
// 	Key   K
// 	Value V
// }

// func (v *KV[K, V]) Set(key K, value V) {
// 	v.Key = key
// 	v.Value = value
// }

// func (v *KV[K, V]) Get(key K) *V {
// 	if v.Key == key {
// 		return &v.Value
// 	}

// 	return nil
// }

// func NewKV[K comparable, V any](key K, value V) *KV[K, V] {
// 	return &KV[K, V]{
// 		Key:   key,
// 		Value: value,
// 	}
// }

// func main() {
// 	var record KV[string, float64]
// 	// var record KV

// 	record.Set("abc", 54.3)

// 	v := record.Get("abc")

// 	if v != nil {
// 		fmt.Println(*v)
// 	}
// }

// func main() {
// 	record := NewKV("abc", 54.3)

// 	v := record.Get("abc")

// 	if v != nil {
// 		fmt.Println(*v)
// 	}

// 	// NewKV("abc", 54.3)
// }

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.05

	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 2.1
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 2.15
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 2.2
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 2.25
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 2.3
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 2.35
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 2.4
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 2.45
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 2.5
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.55
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.6
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.65
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.7
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.75
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.8
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.85
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.9
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 2.95
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.0
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.05
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.1
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.15
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.2
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.25
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.3
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.35
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.4
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.45
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 3.5
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))
}
