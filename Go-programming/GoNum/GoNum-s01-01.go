package main

import (
	"fmt"
	"gonum.org/v1/gonum"
)

func main() {
	version, sum := gonum.Version()

	fmt.Printf("version: %v.\n", version)
	fmt.Printf("sum: %v.\n", sum)
}
