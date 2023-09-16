// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Printf("math.Sqrt(6): %v.\n", math.Sqrt(6))
// }

// package main

// import "fmt"

// type runes []rune

// type myRune rune

// type myString string

// func main() {
// 	// strVarOne := string(runes{0x767d, 0x9d6c, 0x7fd4})

// 	// fmt.Printf("strVarOne: %v.\n", strVarOne)

// 	// strVarOne := string([]myRune{0x266b, 0x266c})

// 	// fmt.Printf("strVarOne: %v.\n", strVarOne)

// 	strVarOne := string([]myRune{0x1f30e})

// 	fmt.Printf("strVarOne: %v.\n", strVarOne)
// }

s := make([]byte, 2, 4)

a0 := [0]byte(s)
a1 := [1]byte(s[1:]) // a1[0] == s[1]
a2 := [2]byte(s)     // a2[0] == s[0]
a4 := [4]byte(s)     // panics: len([4]byte) > len(s)

s0 := (*[0]byte)(s)  // s0 != nil
s1 := (*[1]byte)(s[1:]) // &s1[0] == &s[1]
s2 := (*[2]byte)(s)  // &s2[0] == &s[0]
s4 := (*[4]byte)(s)  // panics: len([4]byte) > len(s)

var t []string
t0 := [0]string(t)  // ok for nil slice t
t1 := (*[0]string)(t)  // t1 == nil
t2 := (*[1]string)(t)  // panics: len([1]string) > len(t)

u := make([]byte, 0)

u0 := (*[0]byte)(u) // u0 != nil

Statement =
	Declaration | LabeledStmt | SimpleStmt |
		GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
		FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt |
		ForStmt | DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assigment | ShortVarDecl .

EmptyStmt = .

LabeledStmt = Label ":" Statement .
Label = identifier .

ExpressionStmt = Expression .

append cap complex imag len make new real
unsafe.Add unsafe.Alignof unsafe.Offsetof unsafe.Sizeof unsafe.Slice

h(x + y)
f.Close()
<-ch
(<-ch)
len("foo")

SendStmt = Channel "<-" Expression .
Channel = Expression .

ch <- 3

IncDecStmt = Expression ( "++" | "--" ) .

x++
x += 1
x--
x -= 1

Assigment = ExpressionList assign_op ExpressionList .
assign_op = [ add_op | mul_op ]	"=" .

x = 1
*p = f()
a[i] = 23
(k) = <-ch

a[i] <<= 2
i &^= 1<<n
