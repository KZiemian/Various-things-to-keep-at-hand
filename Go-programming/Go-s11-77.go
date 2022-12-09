package main

import (
	"fmt"
	"log"
)

func main() {
	name := "Thomas"

	var language = "Go"

	fmt.Println("Hello gophers!")
	log.Println("Hello from log!")

	fmt.Println("Hello", name)

	fmt.Println("We are using", language)
}
