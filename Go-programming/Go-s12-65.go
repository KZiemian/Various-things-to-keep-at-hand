package main

import (
	"fmt"
	"time"
)

// func main() {
// 	c := make(chan string)

// 	go gopher(c)

// 	fmt.Println("--- WAITING")

// 	time.Sleep(2 * time.Second)

// 	fmt.Println("--- WORKING")

// 	for {
// 		v, open := <-c

// 		if open == false {
// 			break
// 		}

// 		fmt.Printf("received: %s.\n", v)
// 	}
// }

// func gopher(c chan string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("sending message")

// 		c <- "gopher"

// 		fmt.Println("message sent")

// 		time.Sleep(500 * time.Millisecond)
// 	}

// 	close(c)
// }

// func main() {
// 	c := make(chan string, 2)

// 	go gopher(c)

// 	fmt.Println("--- WAITING")

// 	time.Sleep(2 * time.Second)

// 	fmt.Println("--- WORKING")

// 	for {
// 		v, open := <-c

// 		if open == false {
// 			break
// 		}

// 		fmt.Printf("received: %v.\n", v)
// 	}
// }

// func gopher(c chan string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("sending message")

// 		c <- "gopher"

// 		fmt.Println("message sent")

// 		time.Sleep(500 * time.Millisecond)
// 	}

// 	close(c)
// }

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go send(200, c1)
	go send(1000, c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func send(sleepMillisecond int, c chan string) {
	for {
		time.Sleep(time.Duration(sleepMillisecond) *
			time.Millisecond)

		c <- fmt.Sprintf("slept %d ms", sleepMillisecond)
	}
}
