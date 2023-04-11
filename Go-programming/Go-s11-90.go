obj, dist := t.scene.Surface.Intersect(ray, infinity)

if obj == nil {
	return t.scene.Env.At(ray.Dir)
}

func (m *Material) absorb(wi, wo Vector3) (bool, Vector3, Vector3)

type Direction Vector3

// Unit normalizes a Vector3 into a Direction.
func (a Vector3) Unit() Direction {
	d := a.Len()
	return Direction{a.X / d, a.Y / d, a.Z / d}
}

type Energy geom.Vector3

func (e Energy) GammaCorrected() Energy {
	return Energy{gamma(e.X), gamma(e.Y), gamma(e.Z)}
}

func (m *Material) absorb(wi, wo Direction) (bool, Vector3, Energy)

type Stringer interface {
	String() string
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

package io

type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	zr, err := gzip.NewReader(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stdin, err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, zr)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

r := io.MultiReader(
	strings.NewReader("<Top>"),
	bytes.NewReader(out),
	strings.NewReader("</Top>"),
)

var logs struct {
	Log []*Commit
}

err = xml.NewDecoder(r).Decode(&logs)

func main() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	ch := make(chan int)

	go fibs(ch)

	for i := 0; i < 20; i++ {
		fmt.Println(<-ch)
	}
}

func fibs(ch chan int) {
	i, j := 0, 1

	for {
		ch <- j
		i, j = j, i + j
	}
}
