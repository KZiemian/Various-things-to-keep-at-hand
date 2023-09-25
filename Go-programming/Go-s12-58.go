package main

import "fmt"

// func main() {
// 	slice := []int{5:1}

// 	fmt.Printf("slice: %v, %T.\n", slice, slice)
// }

// func main() {
// 	var a [2]int

// 	fmt.Printf("a: %v, %v, %T.\n", a, len(a), a)
// }

// func main() {
// 	a := [...]int{1, 2, 3}

// 	fmt.Printf("a: %v, %T.\n", a, a)
// }

// func main() {
// 	a := [...]int{1, 2, 3}
// 	b := a

// 	a[0] = -1
// 	b[1] = -2

// 	fmt.Printf("a: %v.\n", a)
// 	fmt.Printf("b: %v.\n", b)
// }

// func main() {
// 	a := [...]int{1, 2, 3}
// 	b := &a

// 	b[0] = -1
// 	a[1] = -2

// 	fmt.Printf("a: %v, %T.\n", a, a)
// 	fmt.Printf("b: %v, %T.\n", b, b)
// }

// func main() {
// 	line1 := [2]string{"11", "22"}
// 	line2 := [2]string{"21", "22"}
// 	lines := [2][2]string{line1, line2}

// 	fmt.Printf("line1: %v.\n", line1)
// 	fmt.Printf("lines[0]: %v.\n", lines[0])
// 	fmt.Printf("lines[0][0]: %v.\n", lines[0][0])

// }

// func main() {
// 	a := []int{1, 2}

// 	b := a
// 	b[1] = 123

// 	fmt.Printf("a: %v\n", a)
// 	fmt.Printf("b: %v\n", b)
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	b := a[:]

// 	b[1] = -1

// 	fmt.Printf("a: %v.\n", a)
// 	fmt.Printf("b: %v.\n", b)

// 	c := a[4:]

// 	fmt.Printf("c: %v.\n", c)

// 	d := a[:4]

// 	fmt.Printf("d: %v.\n", d)

// 	e := a[2:4]

// 	fmt.Printf("e: %v.\n", e)
// }

// func main() {
// 	a := make([]int, 2)

// 	fmt.Println(a)
// 	fmt.Printf("length: %v\n", len(a))
// 	fmt.Printf("capacity: %v\n", cap(a))

// 	a = append(a, 2)

// 	fmt.Printf("a: %v\n", a)
// 	fmt.Printf("length: %v.\n", len(a))
// 	fmt.Printf("capacity: %v.\n", cap(a))
// }
