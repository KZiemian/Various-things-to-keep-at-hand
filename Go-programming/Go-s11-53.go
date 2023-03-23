package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	var i int = 0
// 	i = 3.2

// 	fmt.Println("i: ", i)
// }

func main() {
	// var i int = 43

	// convertedVar := strconv.Itoa(i)

	// fmt.Printf("convertedVar: %v, %T\n", convertedVar, convertedVar)
	// var f float64 = 41.2

	// convertedVar := strconv.FormatFloat(f, 'f', 4, 64)

	// fmt.Printf("convertedVar: %v, %T\n", convertedVar, convertedVar)

	// var i int64 = 41

	// c1 := strconv.FormatInt(i, 2)
	// fmt.Printf("%v, base 2: %v, %T\n", i, c1, c1)
	// c2 := strconv.FormatInt(i, 8)
	// fmt.Printf("%v, base 8: %v, %T\n", i, c2, c2)
	// c3 := strconv.FormatInt(i, 16)
	// fmt.Printf("%v, base 16: %v, %T\n", i, c3, c3)
	// c4 := strconv.FormatInt(i, 10)
	// fmt.Printf("%v, base 10: %v, %T\n", i, c4, c4)

	// var strVar string = "1"

	// c, err := strconv.ParseFloat(strVar, 64)

	// if err != nil {
	// 	fmt.Println("Something go wrong with strconv.ParseFloat.")
	// }

	// fmt.Printf("c, %v: %T\n", c, c)
	// fmt.Printf("err, %v: %T\n", err, err)

	// var strVar string = "hello"

	// c, err := strconv.ParseFloat(strVar, 64)

	// if err != nil {
	// 	fmt.Println("Something go wrong with strconv.ParseFloat.")
	// }

	// fmt.Printf("c, %v: %T\n", c, c)
	// fmt.Printf("err, %v: %T\n", err, err)

	var strVar string = "101101"

	c, err := strconv.ParseInt(strVar, 2, 64)

	if err != nil {
		fmt.Println("Something go wrong with function strconv.ParseInt")
	}

	fmt.Printf("c, %v: %T\n", c, c)
	fmt.Printf("err, %v: %T\n", err, err)

	strVar = "hello"

	c, err = strconv.ParseInt(strVar, 2, 64)

	if err != nil {
		fmt.Println("Something go wrong with function strconv.ParseInt")
	}

	fmt.Printf("c, %v: %T\n", c, c)
	fmt.Printf("err, %v: %T\n", err, err)
}
