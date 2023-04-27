package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)

		time.Sleep(300 * time.Millisecond)
	}

	quit <- true

	time.Sleep(1000 * time.Millisecond)
}

func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing

			case <-quit:
				fmt.Println("I get enough.")

				return
			}
		}
	}()

	return c
}
