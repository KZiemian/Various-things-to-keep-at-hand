// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func count[T any](slice []T, f func(T) bool) int {
// 	count := 0

// 	for _, s := range slice {
// 		if f(s) {
// 			count++
// 		}
// 	}

// 	return count
// }

// func main() {
// 	s := []string{"ab", "ac", "de", "at", "gfb", "fr"}

// 	fmt.Println(
// 		count(
// 			s,
// 			func(x string) bool {
// 				return strings.Contains(x, "a")
// 			}),
// 	)

// 	s2 := []int{1, 2, 3, 4, 5, 6}

// 	fmt.Println(
// 		count(
// 			s2,
// 			func(x int) bool {
// 				return x % 3 == 0
// 			}),
// 	)
// }

// package main

// import "fmt"

// func unique[T comparable](s []T) []T {
// 	inResult := make(map[T]bool)
// 	var result []T

// 	for _, str := range s {
// 		if _, ok := inResult[str]; !ok {
// 			inResult[str] = true
// 			result = append(result, str)
// 		}
// 	}

// 	return result
// }

// type FruitRank struct {
// 	Fruit string
// 	Rank  uint64
// }

// func main() {
// 	fmt.Println(unique([]string{"abc", "cde", "efg", "efg", "abc",
// 		"cde"}))
// 	fmt.Println(unique([]int{1, 1, 2, 2, 3, 3, 4}))

// 	fruits := []FruitRank{
// 		{
// 			Fruit: "Strawberry",
// 			Rank:  1,
// 		},
// 		{
// 			Fruit: "Rasperry",
// 			Rank:  2,
// 		},
// 		{
// 			Fruit: "Blueberry",
// 			Rank:  3,
// 		},
// 		{
// 			Fruit: "Strawberry",
// 			Rank:  1,
// 		},
// 	}

// 	fmt.Println(unique(fruits))
// }

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 3.55

	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 3.6
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 3.65
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 3.7
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 3.75
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 3.8
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 3.85
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 3.9
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 3.95
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.0
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.05
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.1
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.15
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.2
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.25
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.3
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.35
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.4
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.45
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.5
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.55
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.6
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.65
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.7
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.75
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.8
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.85
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.9
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 4.95
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 5.0
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.05
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.1
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.15
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.2
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.25
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.3
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.35
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.4
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.45
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 5.5
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))
}
