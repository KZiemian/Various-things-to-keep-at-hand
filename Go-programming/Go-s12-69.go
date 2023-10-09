// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func median(data []float64) float64 {
// 	dataCopy := make([]float64, len(data))
// 	copy(dataCopy, data)

// 	sort.Float64s(dataCopy)

// 	var median float64
// 	l := len(dataCopy)

// 	if l == 0 {
// 		return 0
// 	} else if l % 2 == 0 {
// 		median = (dataCopy[l/2 - 1] + dataCopy[l/2]) / 2
// 	} else {
// 		median = dataCopy[l/2]
// 	}

// 	return median
// }

// func main() {
// 	fmt.Println(median([]float64{1, 3, 2}))
// 	fmt.Println(median([]float64{1}))
// 	fmt.Println(median([]float64{3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1,
// 		1, 1, 111}))
// }

// package main

// import (
// 	"fmt"

// 	"golang.org/x/exp/constraints"
// 	"golang.org/x/exp/slices"
// )

// type Number interface {
// 	constraints.Float | constraints.Integer
// }

// func median[T Number](data []T) float64 {
// 	dataCopy := make([]T, len(data))
// 	copy(dataCopy, data)

// 	slices.Sort(dataCopy)

// 	var median float64

// 	l := len(dataCopy)

// 	if l == 0 {
// 		return 0
// 	} else if l % 2 == 0 {
// 		median = float64((dataCopy[l/2 - 1] + dataCopy[l/2]) / 2.0)
// 	} else {
// 		median = float64(dataCopy[l/2])
// 	}

// 	return median
// }

// func main() {
// 	fmt.Printf("median([]float64{1, 3, 2}): %v.\n",
// 		median([]float64{1, 3, 2}))
// 	fmt.Printf("median([]int{1}): %v.\n", median([]int{1}))
// 	fmt.Printf("median([]int32{3, 3, 3, ...}): %v.\n",
// 		median([]int32{3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1, 1,
// 			1, 111}))
// }

// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	myString := "Hello World ✔"

// 	fmt.Printf("len(myString): %v.\n", len(myString))
// 	fmt.Printf("utf8.RuneCountInString(myString): %v.\n",
// 		utf8.RuneCountInString(myString))
// }

// package main

// import "fmt"

// type blog struct {
// 	title  string
// 	author string
// }

// func modifyArray() {
// 	data := []blog{
// 		blog{title: "Title 1", author: "Computer"},
// 		blog{title: "Title 2", author: "Computer"},
// 		blog{title: "Title 3", author: "Computer"},
// 	}

// 	for _, blog := range data {
// 		blog.author = "Supercomputer"
// 	}

// 	fmt.Printf("data: %v.\n", data)
// }

// func modifyArray() {
// 	data := []blog{
// 		blog{title: "Title 1", author: "Computer"},
// 		blog{title: "Title 2", author: "Computer"},
// 		blog{title: "Title 3", author: "Computer"},
// 	}

// 	for idx := range data {
// 		data[idx].author = "Supercomputer"
// 	}

// 	fmt.Printf("data: %v.\n", data)
// }

// func main() {
// 	modifyArray()
// }
