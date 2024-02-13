func merge(out chan<- int, a, b <-chan int) {
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil

				continue
			}

			out <- v

		case v, ok := <-b:
			if !ok {
				b = nil

				contiune
			}

			out <- v
		}
	}

	close(out)
}

type Foo struct {
	f func() error
}

func NewServer(logger func(string, ...interface{})) {
	if logger == nil {
		logger = log.Printf
	}

	logger("initializing %s", os.Getenv("hostname"))
	// ...
}

type Summer interface {
	func Sum() int
}

var t *tree
var s Summer = t

fmt.Println(t == nil, s.Sum())

type ints []int

func (i ints) Sum() int {
	s := 0

	for _, v := range i {
		s += v
	}

	return s
}

func doSum(s Summer) int {
	if s == nil {
		return 0
	}

	return s.Sum()
}

var t *tree
doSum(t)    // (*tree, nil)

var i ints
doSum(i)    // (ints, nil)

doSum(nil)  // (nil, nil)

http.HandleFunc("localhost:8080", nil)

type Writer interface {
	Write(p []byte) (n int, err error)
}

package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("Writing data")
	data := []byte("Some data to write")

	client, _ := net.Dial("tcp", ":3000")
	defer client.Close()

	save(client, data)
}

func save(w io.Writer, data []byte) error {
	_, err := w.Writer(data)

	return err
}
