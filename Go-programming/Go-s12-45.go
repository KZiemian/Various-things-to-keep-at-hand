package main

import "fmt"

func main() {
	// var a [1024]byte
	var s uint = 33

	var i = 1 << s

	fmt.Printf("i: %v.\n", i)

	var j int32 = 1 << s

	fmt.Printf("j: %v.\n", j)
}
