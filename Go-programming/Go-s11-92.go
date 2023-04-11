func main() {
	ch := make(chan int)

	go fibs(ch)

	for i := 0; i < 20; i++ {
		fmt.Println(<-ch)
	}

	ch <- 0
}

func fibs(ch chan int) {
	i, j := 0, 1

	for {
		select {
		case ch <- j:
			i, j = j, i + j

		case <- ch:
			return
		}
	}
}

func main() {
	ch, quit := make(chan int), make(chan int)

	go fibs(ch, quit)

	for i := 0; i < 20; i++ {
		fmt.Println(<-ch)
	}

	quit <- 0 // tell fibs to shut down
}

func fibs(ch, quit chan int) {
	i, j := 0, 1

	for {
		select {
		case ch <- j:
			i, j = j, i + j

		case <-quit:
			return
		}
	}
}

v := doFirstThing()

v1 := doSecondThing(v)

v2 := doThirdThing(v)

doFourthThing(v1, v2)

package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oGtod"))
}
