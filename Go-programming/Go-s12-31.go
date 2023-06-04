func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT%+dh", tz)
}

type List[T any] struct {
	next  *List[T]
	value T
}

// type T[P any] P

// func f[T any]() {
// 	type L T
// }

func (l *List[T]) Len() int {
	// ...
}

TypeParameters = "[" TypeParamList [ "," ] "]" .
TypeParamList  = TypeParamDecl { "," TypeParamDecl } .
TypeParamDecl  = IdentifierList TypeConstriant .

[P any]
[S interface{ ~[]byte | string }]
[S ~[]E, E any]
[P Constraint[int]]
[_ any]

type T[P *C]
type T[P (C)]
type T[P *C | Q]

type T[P interface{*C}]
type T[P *C,]

// type T1[P T1[P]]
// type T2[P interface{ T2[int] }]
// type T3[P interface{ m(T3[int]) }]
// type T4[P T5[P]]
// type T5[P T4[P]]

type T6[P int] struct{ f *T6[P] }

TypeConstraint = TypeElem .

[T []P]
[T ~int]
[T int | string]

int
[]byte
interface{}
interface{ ~int | ~string }
interface{ comparable }
interface{ ~int | ~[]byte }
interface{ ~struct{ any } }

int interface{ ~int }
string comparable
// []byte comparable
any interface{ comparable; int }
any comparable
struct{ f any } comparable
// any interface{ comparable; m() }
interfce{ m() } interface{ comparable; m() }

VarDecl = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
VarSpec = IdentifierList ( Type [ "=" ExpressionList ] | "" ExpressionList ) .

var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]

var d = math.Sin(0.5)  // d is float64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool

ShortVarDecl = IdentifierList ":=" ExpressionList .

"var" IndetifierList "=" ExpressionList .

i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w, _ := os.Pipe()
_, y, _ := coord(p)

field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)
// x, y, x := 1, 2, 3

FunctionDecl = "func" FunctionName [ TypeParameters ] Signature [ FunctionBody ].
FunctionName = identifier .
FunctionBody = Block .
