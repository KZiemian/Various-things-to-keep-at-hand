package main

import (
	"fmt"
	// "sync"
	"time"
)

// var (
// 	count int = 0
// )

// func main() {
// 	wg := sync.WaitGroup{}

// 	wg.Add(1)
// 	go increment(&wg)
// 	wg.Wait()


// 	fmt.Printf("count: %v.\n", count)
// }

// func increment(wg *sync.WaitGroup) {
// 	count++

// 	wg.Done()
// }

// func main() {
// 	wg := sync.WaitGroup{}

// 	wg.Add(1)

// 	go func() {
// 		increment()
// 		wg.Done()
// 	}()

// 	wg.Wait()

// 	fmt.Printf("count: %v.\n", count)
// }

// func increment() {
// 	count++
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	mu := sync.Mutex{}

// 	for i := 0; i < 500; i++ {
// 		wg.Add(1)

// 		go func() {
// 			mu.Lock()
// 			defer mu.Unlock()

// 			increment()

// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()

// 	fmt.Printf("count: %v.\n", count)
// }

// func increment() {
// 	count++
// }

// func main() {
// 	c := make(chan string)

// 	go gopher(c)

// 	v := <-c

// 	fmt.Printf("received: %s.\n", v)
// }

// func gopher(c chan string) {
// 	c <- "gopher"
// }

func main() {
	c := make(chan string)

	go gopher(c)

	for {
		v, open := <-c

		if open == false {
			break
		}

		fmt.Printf("received: %s\n", v)
	}
}

func gopher(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "gopher"
		time.Sleep(500 * time.Millisecond)
	}

	close(c)
}
