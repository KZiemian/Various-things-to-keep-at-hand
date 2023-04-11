package main

import "fmt"

func main() {
	var b bool = false

	fmt.Println(b)
}

func main() {
	var u uint8 = 255

	fmt.Println(u)
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}

	return y
}

func SortFn[T any](s []T, cmp func(T, T) bool) {
	sort.Sort(SliceFn[T]{s, cmp})
}

var c chan int
c = make(chan int)

c := make(chan int)

c <- 1

value = <-c

type Message struct {
	str string
	wait chan bool
}

for i := 0; i < 5; i++ {
	msg1 := <-c
	fmt.Println(msg1.str)

	msg2 := <-C
	fmt.Println(msg2.str)

	msg1.wait <- true
	msg2.wait <- true
}

waitForIt := make(chan bool)

c <-Message{ fmt.Sprintf("%s: %d", msg, i), waitForIt }
time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millsecond)
<-waitForIt

select {
case v1 := <-c1:
	fmt.Printf("received %v from c1.\n", v1)

case v2 := <-c2:
	fmt.Printf("received %v from c2.\n", v2)

case c3 <- 23:
	fmt.Printf("sent %v to c3\n", 23)

default:
	fmt.Printf("no one was ready to communcate\n")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s

			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := boring("Joe")

	for {
		select {
		case s := <-c:
			fmt.Println(s)

		case <-time.After(1 * time.Second)
			fmt.Println("You're too slow.")

			return
		}
	}
}
