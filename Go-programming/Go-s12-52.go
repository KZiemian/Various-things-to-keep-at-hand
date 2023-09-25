// jobID, err := store.PollNextJob()

// if err != nil {
// 	return nil, fmt.Errorf("polling for next job: %w", err)
// }

// owner, err := store.FindOwnerByJobID(jobID)

// if err != nil {
// 	return nil, fmt.Errorf("fetching job owner for job %s: %w", jobID,
// 		err)
// }

// j := jobs.New(jobID, owner)
// res, errr := j.Start()

// if err != nil {
// 	return nil, fmt.Errorf("starting job %s: %w", jobID, err)
// }

// package main

// import "fmt"

// func main() {
// 	sum(1, 2)
// }

// func sum(a, b int) {
// 	fmt.Println(a + b)
// }

// package main

// import "fmt"

// func main() {
// 	r := sum(3, 4)
// 	r = sum(r, 5)

// 	fmt.Println(r)
// }

// func sum(a, b int) int {
// 	return a + b
// }

// package main

// import "fmt"

// func main() {
// 	s, d, p, q := calculate(10, 2)

// 	fmt.Println(s, d, p, q)
// }

// func calculate(a, b int) (int, int, int, int) {
// 	s := a + b
// 	d := a - b
// 	p := a * b
// 	q := a / b

// 	return s, d, p, q
// }

// package main

// import "fmt"

// func main() {
// 	func(a, b int) {
// 		fmt.Println(a + b)
// 	}(10, 2)
// }

// package main

// import "fmt"

// func main() {
// 	sum := func(a, b int) {
// 		fmt.Println(a + b)
// 	}

// 	sum(3, 4)
// }

package main

import "fmt"

func main() {
	sum := func(a, b int) int {
		return a + b
	}

	r := sum(7, 9)

	fmt.Println("r =", r)
}
