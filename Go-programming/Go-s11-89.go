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
