// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}

// 	for x := range Filter(nums, isOdd) {
// 		fmt.Println(x)
// 	}
// }

// func isOdd(n int) bool {
// 	return n % 2 == 1
// }

// package main

// func main() {
// 	const numEvents = 5

// 	events := generateEvents(numEvents)

// 	for _, ev := range events {
// 		process(ev)
// 	}
// }

// package main

// import "sync"

// func main() {
// 	const numEvents = 5

// 	events := generateEvents(numEvents)

// 	var wg sync.WaitGroup

// 	for _, ev := range events {
// 		wg.Add(1)
// 		go process(ev)
// 	}
// }

// package main

// import "sync"

// func main() {
// 	const numEvents = 5

// 	events := generateEvents(numEvents)

// 	var wg sync.WaitGroup

// 	for _, ev := range events {
// 		wg.Add(1)

// 		go func() {
// 			defer wg.Done()
// 			process(ev)
// 		}()
// 	}

// 	wg.Wait()
// }

package main

import (
	"context"
	"sync"
)

func Parallel(events []Event) func(func(int, Event) bool) {
	return func(yield func(int, Event) bool) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var wg sync.WaitGroup
		wg.Add(len(events))

		for idx, event := range events {
			go func() {
				defer wg.Done()

				select {
				case <-ctx.Done():
					return

				default:
					if !yield(idx, event) {
						cancel()
					}
				}

				yield(idx, event)
			}()
		}

		wg.Wait()
	}
}
