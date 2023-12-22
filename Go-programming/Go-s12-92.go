package main

import "fmt"

func main() {
	// var sliceInt []int

	// fmt.Printf("sliceInt: %v, %T.\n", sliceInt, sliceInt)
	// fmt.Printf("len(sliceInt) = %v.\n", len(sliceInt))
	// fmt.Printf("cap(sliceInt) = %v.\n", cap(sliceInt))

	// var sliceInt []int

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("len: %2d, cap: %2d.\n", len(sliceInt),
	// 		cap(sliceInt))

	// 	sliceInt = append(sliceInt, i)
	// }

	// fmt.Printf("sliceInt: %v, %T.\n", sliceInt, sliceInt)

	var m map[int]string

	fmt.Printf("len(m) = %v.\n", len(m))
}

NewGet(
	"http://google.com",
	map[string]string{
		"USER_AGENT": "golang/gopher"
	},
)

// NewGet("http://google.com", map[string]string{})
NewGet("http://google.com", nil)

// func merge(out chan<- int, a, b <-chan int) {
// 	for {
// 		select {
// 		case v := <-a:
// 			out <- v

// 		case v := <-b:
// 			out <-v
// 		}
// 	}
// }

func merge(out chan<- int, a, b <-chan int) {
	var aClosed, bClosed bool

	for !aClosed || !bClosed {
		select {
		case v, ok := <-a:
			if !ok {
				aClosed = true

				continue
			}

			out <- v

		case v, ok := <-b:
			if !ok {
				bClosed = true

				continue
			}
			out <- v
		}
	}

	close(out)
}

func merge(out chan<- int, a, b <-chan int) {
	var aClosed, bClosed bool

	for !aClosed || !bClosed {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil

				continue
			}
			out <- v

		case v, ok := <-b:
			if !ok {
				b = nil

				continue
			}

			out <- v
		}
	}
}
