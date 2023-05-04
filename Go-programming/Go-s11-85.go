// package main

// import (
// 	"fmt"
// 	"time"
// 	"math/rand"
// )

// func main() {
// 	joe := boring("Joe")
// 	ann := boring("Ann")

// 	for i := 0; i < 5; i++ {
// 		fmt.Println(<-joe)
// 		fmt.Println(<-ann)
// 	}

// 	fmt.Println("You're both boring; I'm leaving.")
// }

// func boring(msg string) <-chan string {
// 	c := make(chan string)

// 	go func() {
// 		for i := 0; ; i++ {
// 			c <- fmt.Sprintf("%s %d", msg, i)

// 			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 		}
// 	}()

// 	return c
// }

// 0x.p1
// 1p-2
// 0x1.5e-2
// 1_.5
// 1._5
// 1.5_e1
// 1.5e_1
// 1.5e1_

// imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
// string_lit = raw_string_lit | interpreted_string_lit .
// raw_string_lit = "`" { unicode_char | newline } "`" .
// interpreted_string_lit = `"` { unicode_value | byte_value } `"` .

// var x interface{}
// var v *T
// x = 42
// x = v

// Type = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
// TypeName = identifier | QualifiedIdent .
// TypeArgs = "[" TypeList [ "," ] "]" .
// TypeList = Type { "," Type } .
// TypeLit = ArrayType | StructType | PointerType | FunctionType |
// InterfaceType
// | SliceType | MapType | ChannelType .
// ArrayType = "[" ArrayLength "]" ElementType .
// ArrayLEngth = Expression .
// ElementType = Type .

// [32]byte
// [2*n] struct { x, y int32 }
// [1000]*float64
// [3][5]int
// [2][2][2]float64
// [2]([2]([2]float64))

// // invalid array types
// type (
// 	T1 [10]T1              // element type of T1 is T1
// 	T2 [10]struct { f T2 } // T2 contains T2 as component of a struct
// 	T3 [10]T4              // T3 contains T3 as component of a struct in
// 	// T4
// 	T4 struct { f T3 }     //
// )

// // valid array types
// type (
// 	T5 [10]*T5              // T5 contains T5 as component of a pointer
// 	T6 [10]func() T6        // T6 contains T6 as component of
// 	// a function type
// 	T7 [10]struct{ f []T7 } // T7 contains T7 as component of a slice in
// 	// a struct
// )

// SliceType = "[" "]" ElementType .

// make([]T, length, capacity)
// make([]int, 50, 100)
// new([100]int)[0:50]

StructType = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName [ TypeArgs ] .
Tag = string_lit .

// An empty struct.
struct {}
