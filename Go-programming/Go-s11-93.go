p// ackage main

// import "fmt"

// func main() {
// 	var lengthOfSubframe float64 = 5.2

// 	lengthOfSubframe += 7 * 66.7
// 	lengthOfSubframe += 6 * 4.7

// 	fmt.Printf("Długość podprzedziału w LTE: %v.\n", lengthOfSubframe)
// }

// interface {
// 	P               // illegal: P is a type parameter
// 	int | ~P        // illegal: P is a type parameter
// 	~int | MyInt    // illegal: the type sets for ~int and MyInt are not
// 	// disjoint (~int includes MyInt)
// 	float32 | Float // overlapping type sets but Float is an interface
// }

// var x Float                    // illegal: Float is not a basic interface

// var x interface{} = Float(nil) // illegal

// type Floatish struct {
// 	f Float                // illegal
// }

// // illegal: Bad may not embed itself
// type Bad interface {
// 	Bad
// }

// // illegal: Bad1 may not embed itself using Bad2
// type Bad1 interface {
// 	Bad2
// }

// type Bad2 interface {
// 	Bad1
// }

// // illegal: Bad3 may not embed a union containing Bad3
// type Bad3 interface {
// 	~int | ~string | Bad3
// }

// // illegal: Bad4 may not embed an array containing Bad4 as element type
// type Bad4 interface {
// 	[10]Bad4
// }

// MapType = "map" "[" KeyType "]" ElementType .
// KeyType = Type .

// map[string]int
// map[*T]struct{ x, y float64 }
// map[string]interface{}

// make(map[string]int)
// make(map[string]int, 100)

// ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .

// chan T          // can be used to send and receive vlues of type T
// chan<- float64  // can only be used to send float64s
// <-chan int      // can only be used to receive ints

chan<- chan int    // same as chan<- (chan int)
chan<- <-chan int  // same as chan<- (<-chan int)
<-chan <-chan int  // same as <-chan (<-chan int)
chan (<-chan int)

make(chan int, 100)

type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)

func f[P any](x P) {
	// ...
}

type Celsius float32
type Kelvin  float32

interface{ int }                        // int
interface{ Celsius | Kelvin }           // float32
interface{ ~chan int }                  // chan int
interface{ ~chan int | ~chan<- int }    // chan<- int
interface{ ~[]*data; String() string }  // []*data
