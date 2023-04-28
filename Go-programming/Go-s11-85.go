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

Type = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
TypeName = identifier | QualifiedIdent .
TypeArgs = "[" TypeList [ "," ] "]" .
TypeList = Type { "," Type } .
TypeLit = ArrayType | StructType | PointerType | FunctionType | InterfaceType
| SliceType | MapType | ChannelType .
