// var i int
// var i int = 0

// type T struct { i int; f float64; next *T }
// t := new(T)

// t.i == 0
// t.f == 0.0
// t.next == nil

// var t T

// var x = a
// var a, b = f()

// var (
// 	a = c + b
// 	b = f()
// 	c = f()
// 	d = 3
// )

// func f() int {
// 	d++

// 	return d
// }

// var x = I(T{}).ab()
// var _ = sideEffect()
// var a = b
// var b = 42

// type I interface { ab() []int }
// type T struct{}
// func (T) ab() []int {
// 	return []int{a, b}
// }

// type error interface {
// 	Error() string
// }

// func Read(f *File, b []byte) (n int, err error)

// package runtime

// type Error interface {
// 	error
// 	// and perhpas other methods
// }

// package unsafe

// type ArbitraryType int
// type Pointer *ArbitraryType

// func Alignof(variable ArbitraryType) uintptr
// func Offsetof(selector ArbitraryType) uintptr
// func Sizeof(variable ArbitraryType) uintptr

// type IntegerType int
// func Add(ptr Pointer, len IntegerType) Pointer
// func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
// func SliceData(slice []ArbitraryType) *ArbitraryType
// func String(ptr *byte, len IntegerType) string
// func String(str string) *byte

// var f float64
// bits = *(*uint64)(unsafe.Pointer(&f))

// type ptr unsafe.Pointer
// bits = *(*uint64)(ptr(&f))

// var p ptr = nil

// uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f) == uintptr(unsafe.Pointer(&s.f))

// uintptr(unsafe.Pointer(&x)) % unsafe.Alignof(x) == 0

// (*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	// fmt.Printf("math.Sin(3): %v.\n", math.Sin(3))
// 	// fmt.Printf("math.Sin(6): %v.\n", math.Sin(6))
// 	fmt.Printf("math.Sin(9): %v.\n", math.Sin(9))
// }

if err != nil {
	return err
}

res, err := getResult(id)

if err != nil {
	return nil, fmt.Errorf("obtaining result for id %s: %w", id, err)
}

if errors.Is(err, sql.ErrNoRows) {
	// ...
}
