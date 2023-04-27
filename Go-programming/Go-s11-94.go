package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := boring("Joe")

	for {
		select {
		case s := <-c:
			fmt.Println(s)

		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")

			return
		}
	}
}


func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)

			time.Sleep(time.Duration(rand.Intn(1e3 + 10)) * time.Millisecond)
		}
	}()

	return c
}
