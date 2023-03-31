package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Point []int32

func (p Point) String() string {
	s := ""
	for _, v := range p {
		s += fmt.Sprintf("%v, ", v)
	}

	return s
}

func Scale[S ~[]E, E constraints.Integer](s S, sc E) S {
	r := make(S, len(s))

	for i, v := range s {
		r[i] = v * sc
	}

	return r
}

func ScaleAndPrint(p Point) {
	r := Scale(p, 2)

	fmt.Println(r.String())
}

func main() {
	var pointVar Point = Point([]int32{0, 1, 2, 3})

	ScaleAndPrint(pointVar)
}
