a[low : high : max]

a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]

x.(T)

var x interface{} = 7
i := x.(int)

type I interface { m() }

func f(y I) {
	// s := y.(string)
	r := y.(io.Reader)
}

v, ok = x.(T)
v, ok := x.(T)
var v, ok = ok.(T)
var v, ok interface{} = x.(T)

f(a1, a2, ... an)

math.Atan2(x, y)
var pt *Point
pt.Scale(3.5)

func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

func Join(s, t string) string {
	return s + t
}

if Join(Split(value, len(value)/2)) != value {
	log.Panic("test fails")
}

var p Point
p.Scale(3.5)

func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")

s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)

[P any] // int
[S ~[]E, E any] // []int, int
[P io.Writer]
[P comparable]

func min[T ~int | ~float64](x, y T) T {
	// ...
}

minInt := min[int]
a := minInt(2, 3)
b := min[flaot64](2.0, 3)
c := min(b, -1)

func apply[S ~[]E, E any](s S, f func(E) E) S {
	// ...
}

// f0 := apply[]
f1 := apply[[]int]
f2 := apply[[]string, string]

var bytes []byte
r := apply(bytes, func(byte) byte { ... })

type Vector []float64

func scale[Number ~int64 | ~float64 | ~complex128](v []Number, s Number) []Number

var vector []float64
scaledVector := scale(vector, 42)

func main[T ~int | ~float64](x, y T) T

var x int
min(x, 2.0)
min(1.0, 2.0)

[List ~[]Elem, Elem any]

type Bytes []byte

[A any, B []C, C *A]

Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr = PrimaryExpr | unary_op UnaryExpr .

binary_op = "||" | "&&" | rel_op | add_op | mul_op .
rel_op = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op = "+"| "-" | "|" | "^" .
mul_op = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .
