package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strVar := "13"

	// c, err := strconv.Atoi(strVar)

	// if err != nil {
	// 	fmt.Println("Something go wrong with function strconv.Atoi")
	// }

	// fmt.Printf("c, %v: %T\n", c, c)
	// fmt.Printf("err, %v: %T\n", err, err)

	// strVar = "hello"

	// c, err = strconv.Atoi(strVar)

	// if err != nil {
	// 	fmt.Println("Something go wrong with function strconv.Atoi")
	// }

	// fmt.Printf("c, %v: %T\n", c, c)
	// fmt.Printf("err, %v: %T\n", err, err)

	boolVar := true

	c := strconv.FormatBool(boolVar)

	// if err != nil {
	// 	fmt.Println("Something go wrong with function strconv.FormatBool")
	// }

	fmt.Printf("c, %v: %T.\n", c, c)
	// fmt.Printf("err, %v: %T.\n", err, err)



	// c, err = strconv.FormatBool("hello")

	// if err != nil {
	// 	fmt.Println("Something go wrong with function strconv.FormatBool")
	// }

	// fmt.Printf("c, %v: %T.\n", c, c)
	// fmt.Printf("err, %v: %T.\n", err, err)
}
