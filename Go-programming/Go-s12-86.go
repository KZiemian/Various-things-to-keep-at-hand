// func WithCities(n int) func(*Config) {
// 	// ...
// }

// func main() {
// 	t := NewTerrain(WithCities(9))
// 	// ...
// }

// package math

// func Min(a, b float64) flota64

// package bytes

// func NewReader(b []byte) *Reader

// func WithCities(n int) func(*Config)

// type Option interface {
// 	Apply(*Config)
// }

// func NewTerrain(options ...Option) *Terrain {
// 	var config Config

// 	for _, option := range options {
// 		option.Apply(&config)
// 	}

// 	// ...
// }

// type splines struct {
// 	// ...
// }

// func (s *splines) Apply(c *Config) {
// 	// ...
// }

// func WithReticulatedSplines() Option {
// 	return new(splines)
// }

// type cities struct {
// 	cities int
// }

// func (c *cities) Apply(c *Config) {
// 	// ...
// }

// func WithCities(n int) Option {
// 	return &cities{
// 		cities: n,
// 	}
// }

// func main() {
// 	t := NewTerrain(
// 		WithRecticulatedSplines(),
// 		WithCities(9),
// 	)
// }

package main

import "fmt"

type Calculator struct {
	acc float64
}

const (
	OP_ADD = 1 << iota
	OP_SUB
	OP_MUL
)

func (c *Calculator) Do(op int, v float64) float64 {
	switch op {
	case OP_ADD:
		c.acc += v

	case OP_SUB:
		c.acc -= v

	case OP_MUL:
		c.acc *= v

	default:
		panic("unhadled operation")
	}

	return c.acc
}

func main() {
	var c Calculator

	fmt.Printf("c.Do(OP_ADD, 100): %v.\n", c.Do(OP_ADD, 100))
	fmt.Printf("c.Do(OP_SUB, 50):  %v.\n", c.Do(OP_SUB, 50))
	fmt.Printf("c.Do(OP_MUL, 3):   %v.\n", c.Do(OP_MUL, 3))
}
