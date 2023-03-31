package main

import "fmt"

type Tree[T interface{}] struct {
	left, right *Tree[T]
	data
}

func (t *Tree[T]) Lookup(x T) *Tree[T]

var stringTree Tree[string]

func min(x, y float64) float64

func min[T constraints.Ordered](x, y T) T

type Ordered interface {
	Integer|Float|~string
}

[S interface{~[]E}, E interface{}]

[S ~[]E, E interface{}]

[S ~[]E, E any]

func min[T constraints.Ordered](x, y T) T

var a, b, m float64

m = min[float64](a, b)

m = min(a, b)

func Scale[E constraints.Integer](s []E, sc E) []E {
	r := make([]E, len(s))

	for i, v := range s {
		r[i] = v * sc
	}

	return r
}

type Point []int32

func (p Point) String() string {
	// ...
}

func ScaleAndPrint(p Point) {
	r := Scale(p, 2)

	fmt.Println(r.String())
}

func Scale[S ~[]E, E constraints.Integer](s S, sc E) S {
	r := make(S, len(s))

	for i, v := range s {
		r[i] = v * sc
	}

	return r
}

type Tree[T any] struct {
	cmp  func(T, T) int
	root *node[T]
}

type node[T any] struct {
	left, right *node[T]
	data        T
}

func (bt *Tree[T]) find(val T) **node[T] {
	pl := &bt.root

	for *pl != nil {
		switch cmp := bt.cmp(val, (*pl).data); {
		case cmp < 0:
			pl = &(*pl).left
		case cmp > 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}

	return pl
}

type SliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.s)
}

func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func (s SliceFn[T]) Less(i, j int) bool {
	return s.cmp(s.s[i], s.s[j])
}
