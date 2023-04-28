// package main

// import "fmt"

// func main() {
// 	fmt.Printf("Hello, Vlad!")
// }

// package main

// import "fmt"

// func main() {
// 	suma := 0.0
// 	// suma = 167.05 + 25.67 + 41.93 + 167.05 + 111.25 + 28.58 + 132.93

// 	// fmt.Printf("Suma składek za luty: %v PLN.\n", suma)

// 	// suma = 417.63 + 64.19 + 104.84 + 332.31 + 417.63 + 278.14 + 71.46

// 	// fmt.Printf("Suma składek za marzec: %v PLN.\n", suma)

// 	suma += 468.48 + 72.00 + 117.60 + 372.77 + 468.48 + 312.00 + 80.16

// 	fmt.Printf("Suma składek za czerwiec: %v PLN.\n", suma)
// }

// Syntax = { Production } .
// Production = prodution_name "=" [ Expression ] "." .
// Expression = Term { "|" Term } .
// Factor = production_name | token [ "..." token ] | Group | Option | Repetiton .
// Group = "(" Expression ")" .
// Option = "[" Expression "]" .
// Repetiton = "{" Expression "}" .

// newline = /* the Unicode code point U+000A */ .
// unicode_char = /* an arbitrary Unicode code point except newline */ .
// unicode_letter = /* a Unicode code point categorized as "Letter" */ .
// unicode_digit = /* a Unicode code point categoried as "Number, decimal digit" */ .

// letter = unicode_letter | "_" .
// decimal_digit = "0" ... "9" .
// binary_digit = "0" | "1" .
// octal_digit = "0" ... "7" .
// hex_digit = "0" ... "9" | "A" ... "F" | "a" ... "f" .

// indentifier = letter { letter | unicode_digit } .
// a
// _x9
// ThisVariableIsExported

// break
// case
// chan
// const
// continue
// default
// defer
// else
// fallthrough
// for
// func
// go
// goto
// if
// import
// interface
// map
// package
// range
// return
// select
// struct
// switch
// type
// var
// +
// -
// *
// /
// %
// &
// |
// ^
// <<
// >>
// &^
// +=
// -=
// *=
// /=
// %=
// &=
// |=
// ^=
// <<=
// >>=
// &^=
// &&
// ||
// <-
// ++
// --
// ==
// <
// >
// =
// !
// ~
// !=
// <=
// >=
// :=
// ...
// (
// [
// {
// ,
// .
// )
// ]
// }
// ;
// :
