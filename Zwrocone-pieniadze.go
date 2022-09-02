package main

import "fmt"

func main() {
	zwroconePieniadze := make([]float64, 20)



	// ##############################
	// Rok 2022

	// Game „Rayman Origins”, 22 VIII 2022
	zwroconePieniadze[0] = 39.90



	sum := 0.0

	for _, v := range zwroconePieniadze {
		sum += v
	}

	fmt.Printf("Za filmy, gry, komiksy, książki, muzykę, zwróciłeś:")
	fmt.Printf(" %v PLN.\n", sum)
}
