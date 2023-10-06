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

package main

import "fmt"

func unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T

	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}

	return result
}

type FruitRank struct {
	Fruit string
	Rank  uint64
}

func main() {
	fmt.Println(unique([]string{"abc", "cde", "efg", "efg", "abc",
		"cde"}))
	fmt.Println(unique([]int{1, 1, 2, 2, 3, 3, 4}))

	fruits := []FruitRank{
		{
			Fruit: "Strawberry",
			Rank:  1,
		},
		{
			Fruit: "Rasperry",
			Rank:  2,
		},
		{
			Fruit: "Blueberry",
			Rank:  3,
		},
		{
			Fruit: "Strawberry",
			Rank:  1,
		},
	}

	fmt.Println(unique(fruits))
}
