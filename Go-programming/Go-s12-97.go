// func main() {
// 	fmt.Println("Writing data")
// 	data := []byte("Some data to write")

// 	client, _ := net.Dial("tcp", ":3000")
// 	defer client.Close()

// 	save(LogWriter(client), data)
// }

// type logWriter struct {
// 	writer io.Writer
// }

// func (l logWriter) Write(p []byte) (n int, err error) {
// 	fmt.Printf("%s\n", p)

// 	return l.writer.Write(p)
// }

// func LogWriter(w io.Writer) io.Writer {
// 	return logWriter{w}
// }

// package main

// import ("fmt"
// 	"math")

// func main() {
// 	// sqrtOf2order12 := 1.0594630943
// 	sqrtOf2order12 := 1.05946309

// 	result := math.Pow(sqrtOf2order12, 6.0)

// 	fmt.Printf("sqrtOf2order12^6 = %v.\n", result)

// 	result *= result

// 	fmt.Printf("sqrtOf2order12^12 = %v.\n", result)
// }

// package main

// import "fmt"

// func backwards(xs []int) func(func(i int, x int) bool) {
// 	return func(yield func(i int, x int) bool) {
// 		for i := len(xs) - 1; i >= 0; i-- {
// 			if !yield(i, xs[i]) {
// 				return
// 			}
// 		}
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}

// 	for _, x := range backwards(nums) {
// 		fmt.Println(x)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}

// 	for x := range Map(nums, toSquare) {
// 		fmt.Println(x)
// 	}
// }

// func toSquare(n int) int {
// 	return n * n
// }

package main

import (
	"strconv"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for x := range Map(nums, strconv.Itoa) {
		fmt.Printf("%s is a %T.\n", x, x)
	}
}

func toSquare(n int) int {
	return n * n
}
