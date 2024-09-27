// package main

// import "fmt"

// func main() {
// 	err := openFile()

// 	if err != nil {
// 		fmt.Println(err.Error())

// 		return
// 	}

// 	defer closeFile()


// 	err = readData()

// 	if err != nil {
// 		fmt.Println(err.Error())

// 		return
// 	}


// 	err = writeData()

// 	if err != nil {
// 		fmt.Println(err.Error())

// 		return
// 	}


// 	fmt.Println("no errors")
// }

// func openFile() error {
// 	fmt.Println("opening file connection")

// 	return nil
// }

// func closeFile() {
// 	fmt.Println("closing file connection")
// }

// func readData() error {
// 	fmt.Println("reading file data")

// 	return nil
// }

// func writeData() error {
// 	fmt.Println("writing file data")

// 	return nil
// }

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.01

	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.05
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.1
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.15
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.2
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 0.25
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 0.3
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 0.35
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 0.4
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 0.45
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 0.5
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))
}
