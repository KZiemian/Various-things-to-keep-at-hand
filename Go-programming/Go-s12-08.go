package main

import (
	"fmt"
	// "math"
)

// type myString string

// type bytes []byte

// type myByte byte

func main() {
	// eApprox := 2.0
	// seriesTerm := 1.0

	// fmt.Printf("2 wyraz szeregu: %v.\n", eApprox)

	// for i := 2; i <= 15; i++ {
	// 	seriesTerm /= float64(i)
	// 	eApprox += seriesTerm

	// 	fmt.Printf("%v wyrazy szeregu: %v.\n", i + 1, eApprox)
	// }

	// fmt.Printf("math.Pow(2, 3): %v.\n", math.Pow(2, 3))

	// fmt.Printf("5 / 3 = %v.\n", 5/3)
	// fmt.Printf("5 %% 3 = %v.\n", 5%3)
	// fmt.Printf("-5 / 3 = %v.\n", -5/3)
	// fmt.Printf("-5 %% 3 = %v.\n", -5%3)
	// fmt.Printf("5 / -3 = %v.\n", 5/-3)
	// fmt.Printf("5 %% -3 = %v.\n", 5 % -3)
	// fmt.Printf("-5 / -3 = %v.\n", -5/-3)
	// fmt.Printf("-5 %% -3 = %v.\n", -5%-3)

	// fmt.Printf("1 / 0 = %v.\n", 1/0)

	// x := -11

	// fmt.Printf("x & 3: %v.\n", x & 3)

	// fmt.Printf("1.0 / 0.0: %v.\n", 1.0 / 0.0)

	// var x float64 = 1.0
	// var zeroValue float64 = 0.0

	// fmt.Printf("x / zeroValue: %v.\n", x / zeroValue)

	// stringVar := string(65)

	// fmt.Printf("stringVar: %v.\n", stringVar)

	// symbolStr := string(0x266c)

	// fmt.Printf("symbolStr: %v.\n", symbolStr)

	// var intVar int = int(1.2)

	// fmt.Printf("intVar: %v.\n", intVar)

	// var strVar string = string(65.0)

	// fmt.Printf("strVar: %v.\n", strVar)

	// var intVar int = 1
	// var uintVar uint = 1

	// uintVar = intVar

	// fmt.Printf("uintVar: %v.\n", uintVar)

	// strVar := string('a')

	// fmt.Printf("string('a'): %v.\n", string('a'))
	// // fmt.Printf("string(-1): %v.\n", string(-1))
	// fmt.Printf("string(0xf8): %v.\n", string(0xf8))

	// fmt.Printf("myString(0x65e5): %v.\n", myString(0x65e5))

	// strVar := string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'})

	// fmt.Printf("strVar: %v.\n", strVar)

	// varOne := string(bytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'})

	// fmt.Printf("varOne: %v.\n", varOne)

	// strVarOne := string([]myByte{'w', 'o', 'r', 'l', 'd', '!'})

	// fmt.Printf("strVarOne: %v.\n", strVarOne)

	// strVarOne := myString([]myByte{'\xf0', '\x9f', '\x8c', '\x8d'})

	// fmt.Printf("strVarOne: %v.\n", strVarOne)

	strVarOne := string([]rune{0x767d, 0x9d6c, 0x7fd4})

	fmt.Printf("strVarOne: %v.\n", strVarOne)
}
