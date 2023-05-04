// func main() {
// 	ch := make(chan int)

// 	go fibs(ch)

// 	for i := 0; i < 20; i++ {
// 		fmt.Println(<-ch)
// 	}

// 	ch <- 0
// }

// func fibs(ch chan int) {
// 	i, j := 0, 1

// 	for {
// 		select {
// 		case ch <- j:
// 			i, j = j, i + j

// 		case <- ch:
// 			return
// 		}
// 	}
// }

// func main() {
// 	ch, quit := make(chan int), make(chan int)

// 	go fibs(ch, quit)

// 	for i := 0; i < 20; i++ {
// 		fmt.Println(<-ch)
// 	}

// 	quit <- 0 // tell fibs to shut down
// }

// func fibs(ch, quit chan int) {
// 	i, j := 0, 1

// 	for {
// 		select {
// 		case ch <- j:
// 			i, j = j, i + j

// 		case <-quit:
// 			return
// 		}
// 	}
// }

// v := doFirstThing()

// v1 := doSecondThing(v)

// v2 := doThirdThing(v)

// doFourthThing(v1, v2)

// package main

// import (
// 	"fmt"

// 	"github.com/golang/example/stringutil"
// )

// func main() {
// 	fmt.Println(stringutil.Reverse("!oGtod"))
// }

// An interface representing all types with underlying type int.
interface {
	~int
}

// An interface representing all types with underlying type int that
// implement the String method.
interface {
	~int
	String() string
}

// An interface representing an empty type set: there is no type that is both
// an int and a string.
interface {
	int
	string
}

type MyInt int

interface {
	~[]byte  // the underlying type of []byte is itself
	~MyInt   // illegal: the underlying type of MyInt is not MyInt
	~error   // illegal: error is an interface
}

// The Float interface represents all floating-point types
// (including any named types whose underlying types are
// either flaot32 or float64).
type Float interface {
	~float32 | ~float64
}
