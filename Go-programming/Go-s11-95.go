package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var randNum int

	for i := 0; i < 200; i++ {
		randNum = rand.Intn(1e3 + 10)

		if randNum >= 1000 {
			fmt.Printf("\nUwaga!\n")
			fmt.Printf("randNum: %v.\n\n", randNum)
		} else {
			fmt.Printf("randNum: %v, %T.\n", randNum, randNum)
		}


		time.Sleep(300 * time.Millisecond)
	}
}
