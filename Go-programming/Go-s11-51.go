package main

import "fmt"

func main() {
	a := 3.0
	b := 6.0

	if a == 0.0 {
		if b == 0.0 {
			fmt.Println("Tożsamość.")
		} else {
			fmt.Println("Sprzeczność.")
		}
	} else {
		fmt.Printf("Rozwiązanie: %v.\n", -b / a)
	}
}
