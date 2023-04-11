package main

import "fmt"

func main() {
	a := 2
	b := 5

	var1, var2, var3, var4 := calculateArithmetic(a, b)

	fmt.Printf("%v + %v = %v.\n", a, b, var1)
	fmt.Printf("%v - %v = %v.\n", a, b, var2)
	fmt.Printf("%v * %v = %v.\n", a, b, var3)
	fmt.Printf("%v / %v = %v.\n", a, b, var4)
}

func calculateArithmetic(a, b int) (int, int, int, int) {
	s := a + b
	d := a - b
	p := a * b
	q := a / b

	return s, d, p, q
}
