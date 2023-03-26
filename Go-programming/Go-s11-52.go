package main

import "fmt"

func main() {
	var b bool = false

	fmt.Println(b)
}

func main() {
	var u uint8 = 255

	fmt.Println(u)
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}

	return y
}
