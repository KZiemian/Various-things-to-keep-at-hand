package main

import (
	"fmt"
)

func main() {
	// func(a, b int) {
	// 	fmt.Printf("%v + %v = %v.\n", a, b, a + b)
	// }(10, 3)

	sumFun := func(a, b int) {
		fmt.Printf("%v + %v = %v.\n", a, b, a + b)
	}

	sumFun(1, 2)
}
