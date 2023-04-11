package main

import "fmt"

func main() {
	var lengthOfSubframe float64 = 5.2

	lengthOfSubframe += 7 * 66.7
	lengthOfSubframe += 6 * 4.7

	fmt.Printf("Długość podprzedziału w LTE: %v.\n", lengthOfSubframe)
}
