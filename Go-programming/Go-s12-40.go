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
