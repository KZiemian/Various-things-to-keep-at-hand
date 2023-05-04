// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// func(a, b int) {
// 	// 	fmt.Printf("%v + %v = %v.\n", a, b, a + b)
// 	// }(10, 3)

// 	sumFun := func(a, b int) {
// 		fmt.Printf("%v + %v = %v.\n", a, b, a + b)
// 	}

// 	sumFun(1, 2)
// }

// float_lit = decimal_float_lit | hex_float_lit .
// 	decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponents ] | decimal_digits decimal_eponent |
// 	"." decimal_digits [ decimal_exponent ] .
// decimal_exponent = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

// hex_float_lit = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
// hex_mantissa = [ "_" ] hex_digits "." [ hex_digits ] |
// [ "_" ] hex_digits |
// "." hex_digits.
// hex_exponent = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .

// rune_lit = "'" ( unicode_value | byte_value ) "'" .
// unicode_value = unicode_char | little_u_value | big_u_value | escaped_char .
// byte_value = octal_byte_value | hex_byte_value .
// octal_byte_value = `\` octal_digit octal_digit octal_digit .
// hex_byte_value = `\` "x" hex_digit hex_digit .
// little_u_value = `\` "u" hex_digit hex_digit hex_digit .
// big_u_value = `\` "U" hex_digit hex_digit hex_digit
// hex_digit hex_digit hex_digit .
// 	escaped_char = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | '"' ) .

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}

// A struct with four embedded fields of types T1, *T2, P.T3 and *P.T4.
struct {
	T1       // field name is T1
	*T2      // field name is T2
	P.T3     // field name is T3
	*P.T4    // field name is T4
	x, y int // field names are x and y
}

struct {
	T    // conflicts with embedded field *T and *P.T
	*T   // conflicts with embedded field T and *P.T
	*P.T // conflicts with embedded field T and *T
}

struct {
	x, y float64 "" // an empty tag string is like an absent tag
	name string "any string is permitted as a tag"
	_    [4]byte "ceci n'est pas un champ de structure"
}

// A struct corresponding to a TimeStamp protocol buffer.
// The tag strings define the protocol buffer field numbers;
// they follow the convention outlined by the reflect package.
struct {
	microsec  uint64 `protobuf:"1"`
	serverIP6 uint64 `protobuf:"2"`
}

// invalid struct types
type (
	T1 struct{ T1 }       // T1 contains a field of T1
	T2 struct{ f [10]T2 } // T2 contains T2 as component of an array
	T3 struct{ T4 }       // T3 contains T3 as component of array
	// in struct T4
	T4 struct{ f [10]T3 } // T4 contains T4 as component of struct T3
	// in an array
)

// valid struct types
type (
	T5 struct{ f *T5 }       // T5 contains T5 as component of a pointer
	T6 struct{ f func() T6 } // T6 contains T6 as component of a function
	// type
	T7 struct{ f [10][]T7 }  // T7 contains T7 as component of a slice
	// in an array
)

PointerType = "*" BaseType .
BaseType    = Type .

*Point
*[4]int
