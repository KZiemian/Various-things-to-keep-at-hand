// interface{}                           // no single underlying type
// interface{ Celsius | float64 }        // no single underlying type
// interface{ chan int | chan<- string } // channels have different element types
// interface{ <-chan int | chan<- int }  // directional channels have different
// // directions

// interface{ int }                 // int (same as ordinary core type)
// interface{ []byte | string }     // bytestring
// interface{ ~[]byte | myString }  // bytestring

// type (
// 	A0 = []string
// 	A1 = A0
// 	A2 = struct{ a, b int }
// 	A3 = int
// 	A4 = func(A3, float64) *A0
// 	A5 = func(x int, _ float64) *[]string

// 	B0 A0
// 	B1 []string
// 	B2 struct{ a, b int }
// 	B3 struct{ a, c int }
// 	B4 func(int, float64) *B0
// 	B5 func(x int, y float64) *A1

// 	C0 = B0
// 	D0[P1, P2 any] struct{ x P1; y P2 }
// 	E0 = D0[int, string]
// )

// // These types are identical.
// A0, A1, and []string
// A2 and struct{ a, b int }
// A3 and int
// A4, func(int, float64) *[]string, and A5

// B0 and C0
// D0[int, string] and E0
// []int and []int
// struct{ a, b *B5 } and struct{ a, b *B5 }
// func(x int, y float64) *[]string, func(int, float64) (result *[]string),
// and A5

// x T x is representable by a value of T because

// 'a' byte 97 is in the set of byte values
// 97 rune rune is an alias for int32, and 97 is in set of 32-bit integers
// "foo" string "foo" is in the set of string values
// 1024 int16 1024 is in the set of 16-bit integers
// 42.0 byte 42 is in the set of unsigned 8-bit integers
// 1e10 uint64 10000000000 is in the set of unsigned 64 bit integers
// 2.718281828459045 float32 2.718281828459045 rounds to 2.7182817 which is in the set of float32 values
// -1e-1000 float64 -1e-1000 rounds to IEEE -0.0 which is further simplified to 0.0
// 0i int 0 is an integer value
// (42 + 0i) float32 42.0 (with zero imaginary part) is in the set of float32 values

// x T x is not representable by a value of T because
// 0 bool 0 is not in the set of boolean values
// 'a' string 'a' is a rune, it is not in the set of string values
// 1024 byte 1024 is not in the set of unsigned 8-bit integers
// -1 uint16 -1 is not in the set of unsigned 16-bit integers
// 1.1 int 1.1 is not an integer value
// 42i flaot32 (0 + 42i) is not in the set of float32 values
// 1e1000 float64 1e1000 overflows to IEEE +Inf after rounding

// Block = "{" StatementList "}" .
// StatementList = { Statement ";" } .

// Declaration = ConstDecl | TypeDecl | VarDecl .
// TopLEvelDecl = Declaration | FunctionDecl | MethodDecl .

// Types:
// any bool byte comparable
// complex64 complex128 error float32 float64
// int int8 int16 int64 rune string
// uint uint8 uint16 uint32 uint64 uintptr

// Constants:
// true false iota

// Zero value:
// nil

// Functions:
// append cap close complex copy delete imag len
// make new panic print println real recover

ConstDecl = "const" ( ConstSpec | "(" { ConstSpec ";" } ")" .
ConstSpec = IdetifierList [ [ Type ] "=" ExpressionList ] .

IdentifierList = identifier { "," identifier } .
ExpressionList = Expression { "," Expression } .

const Pi float64 = 3.14159265358979323846
const zero = 0.0    // untyped floating-point constant

const (
	size int64 = 1024
	eof        = -1  // untyped integer constant
)
