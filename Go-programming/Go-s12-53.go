// package main

// import "fmt"

// func main() {
// 	fmt.Printf("146 %% 25 = %v.\n", 146 % 25)
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println(calculate(1, 5, sum))
// }

// func calculate(a, b int, f func(a, b int) int) int {
// 	return f(a, b)
// }

// func sum(a, b int) int {
// 	return a + b
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println(calculate(1, 5, sum))
// 	fmt.Println(calculate(2, 1, substract))
// 	fmt.Println(calculate(10, 2, func(a, b int) int {
// 		return a * b
// 	}))
// }

// func calculate(a, b int, f func(a, b int) int) int {
// 	return f(a, b)
// }

// func sum(a, b int) int {
// 	a + b
// }

// func subtract(a, b int) int {
// 	return a - b
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now().UnixMilli()

// 	t := timer(now)

// 	time.Sleep(200 * time.Millisecond)

// 	fmt.Println(t())
// }

// func timer(start int64) func() int64 {
// 	return func() int64 {
// 		end := time.Now().UnixMilli()

// 		return end - start
// 	}
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := timer()

	time.Sleep(200 * time.Millisecond)

	t2 := timer()

	time.Sleep(200 * time.Millisecond)

	fmt.Println(t1())
	fmt.Println(t2())
}

func timer() func() int64 {
	start := time.Now().UnixMilli()

	return func() int64 {
		end := time.Now().UnixMilli()

		return end - start
	}
}
