// package main

// import (
// 	"fmt"
	// "math"

	// "golang.org/x/exp/constraints"
// )

// func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
// 	var newValues []T

// 	for _, v := range values {
// 		newValue := mapFunc(v)

// 		newValues = append(newValues, newValue)
// 	}

// 	return newValues
// }

// func main() {
// 	result1 := MapValues([]int{1, 2, 3}, func(n int) int {
// 		return 2 * n
// 	})

// 	fmt.Printf("result1: %v.\n", result1)

// 	result2 := MapValues([]float64{1.0, 2.0, 3.0}, math.Sin)

// 	fmt.Printf("result2: %v.\n", result2)
// }

// type CustomData interface {
// 	constraints.Ordered | []byte | []rune
// }

// type User[T CustomData] struct {
// 	ID   int
// 	Name string
// 	Data T
// }

// func main() {
// 	u := User[int]{
// 		ID: 0,
// 		Name: "Jason",
// 		Data: 3,
// 	}

// 	fmt.Printf("user: %v.\n", u)

// 	uAnother := User[string]{
// 		ID: 1,
// 		Name: "Rob",
// 		Data: "interface",
// 	}

// 	fmt.Printf("uAnother: %v.\n", uAnother)
// }

// type CustomMap[T comparable, V int | string] map[T]V

// func main() {
// 	// m := make(CustomMap[int, string])
// 	// m[3] = "three"

// 	// fmt.Printf("m: %T, %v.\n", m, m)

// 	// fmt.Printf("2 * math.Sqrt2 * math.Pi^4 * sqrt[4](math.Pi): %v.\n",
// 	// 	math.Pow(8, 0.25) * math.Pow(math.Pi, 2) *
// 	// 		math.Pow(math.Pi, 0.125))
// 	gamma := 1.01

// 	result := 0.5 / (gamma * (gamma + 1) * (gamma - 1))

// 	fmt.Printf("result: %v.\n", result)
// }

package main

import (
	"testing"
)

func TestSomething(t *testing.T) {
	t.Error()
}

var r io.Reader
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)

var r io.Reader
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

if err != nil {
	return nil, err
}

r = tty
