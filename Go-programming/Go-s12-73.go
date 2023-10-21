// var w io.Writer
// w = r.(io.Writer)

// var empty interface{}
// empty = w

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	constForChi := 8 * math.SqrtPi * math.Pow(math.Pi, 8.0)
// 	alpha := 19.1
// 	// fmt.Printf("1.0/14.0: %v.\n", 1.0/14.0)

// 	boundForLambda := constForChi / math.Pow(alpha, 4.0)

// 	fmt.Printf("boundForLambda: %v.\n", boundForLambda)

// 	alpha1 := 2 * math.Pi * math.Pi * math.Pow(math.Pi, 0.125)
// 	fmt.Printf("alpha1: %v.\n", alpha1)

// 	alpha = 22.78
// 	boundForLambda = constForChi / math.Pow(alpha, 4.0)
// 	fmt.Printf("boundForLambda: %v.\n", boundForLambda)

// 	alpha2 := math.Pow(80, 0.25) * math.Pi * (math.Pi *
// 		math.Pow(math.Pi, 0.125))
// 	fmt.Printf("alpha2: %v.\n", alpha2)

// 	alpha = 34.06

// 	boundForLambda = constForChi / math.Pow(alpha, 4.0)

// 	fmt.Printf("boundForLambda: %v.\n", boundForLambda)

// 	alpha3 := math.Pow(80_000, 0.25) * math.Pi * (math.Pi *
// 		math.Pow(math.Pi, 0.125))

// 	fmt.Printf("alpha3: %v.\n", alpha3)

// 	boundForLambda = constForChi / math.Pow(alpha3, 4.0)

// 	fmt.Printf("boundForLambda: %v.\n", boundForLambda)
// }

package main

import "fmt"

// func printAnything[T any](thing T) {
// 	fmt.Println(thing)
// }

// func main() {
// 	printAnything("Hello!")
// 	printAnything(42)
// 	printAnything(true)
// }

// func Join[E any](things []E) (result string) {
// 	for _, v := range things {
// 		result += v.String()
// 	}

// 	return result
// }

// type Stringer interface {
// 	String() string
// }

// func Join[E Stringer](things []E) (result string) {
// 	for _, v := range things {
// 		result += v.String()
// 	}

// 	return result
// }

// func main() {
// 	output := Join([]string{"a", "b", "c"})

// 	fmt.Printf("output: %v.\n", output)
// }

// func Equal[T any](a, b T) bool {
// 	return a == b
// }

// func Equal[T comparable](a, b T) bool {
// 	return a == b
// }

// func main() {
// 	fmt.Printf("Equal(1, 1): %v.\n", Equal(1, 1))
// }
