package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := boring("Joe")

	timeout := time.After(5 * time.Second)

	for {
		select {
		case s := <-c:
			fmt.Println(s)

		case <-timeout:
			fmt.Println("You talki too much.")

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
