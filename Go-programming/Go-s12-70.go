package main

import "fmt"

// func main() {
// 	a1 := []int{1, 2, 3}

// 	fmt.Printf("a1: %v.\n", a1)

// 	a2 := a1[1:]
// 	fmt.Printf("a2: %v.\n", a2)

// 	for i := range a2 {
// 		a2[i] += 10
// 	}

// 	fmt.Printf("a1: %v.\n", a1)
// 	fmt.Printf("a2: %v.\n", a2)

// 	a2 = append(a2, 4)
// 	fmt.Printf("a1: %v.\n", a1)
// 	fmt.Printf("a2: %v.\n", a2)

// 	for i := range a2 {
// 		a2[i] += 10
// 	}

// 	fmt.Printf("a1: %v.\n", a1)
// 	fmt.Printf("a2: %v.\n", a2)
// }

func testInterface() {
	var data *byte
	var testInterface interface{}

	fmt.Printf("data == nil: %v.\n", data == nil)
	fmt.Printf("testInterface == nil: %v.\n", testInterface == nil)

	testInterface = data

	fmt.Printf("testInterface == nil: %v.\n", testInterface == nil)
}

func main() {
	testInterface()
}
