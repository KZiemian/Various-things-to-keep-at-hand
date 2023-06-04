// func IndexRune(s string, r rune) int {
// 	for i, c := range s {
// 		if c == r {
// 			return i
// 		}
// 	}
// 	// invalid: missing return statement
// }

func min[T ~int | ~float64](x, y T) T {
	if x < y {
		return x
	}

	return y
}

func flushICache(begin, end uintprt) // implemented externally

MethodDecl = "func" Receiver MethodName Signature [ FunctionBody ] .
Receiver = Parameters .

func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(fator float64) {
	p.x *= factor
	p.y *= factor
}

type Pair[A, B any] struct {
	a A
	b B
}

func (p Pari[A, B]) Swap() Pair[B, A] {
	// ...
}

func (p Pair[First, _]) First() First {
	// ...
}

Operand = Literal | OperangName [ TypeArgs ] | "(" Expression ")" .
Literal = BasicLit | CompositeLit | FunctionLit .
BasicLit = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
OperandName = identifier | QualifiedIdent .

QulifiedIdent = PackageName "." identifier .

math.Sin

CompositeLit = LiteralType LiteralValue .
LiteralType = StructType | ArrayType | "[" "..." "]" ElementType |
SliceType | MapType | TypeName [ TypeArgs ] .
LiteralValue = "{" [ ElementList [ "," ] ] "}" .
ElementList = KeyedElement { "," KeyedElment } .
KeyedElement = [ Key ":" ] Element .
Key = FieldName | Expression | LiteralValue .
FieldName = identifier .
Element = Expression | LiteralValue .

type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }

origin := Point3D{}
line := Line{origin, Point3D{y: -4, z: 12.3}}

var pointer *Point3D = &Point3D{y: 1000}

p1 := &[]int{}
p2 := new([]int)

buffer := [10]string{}  // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}  // len(intSet) == 6
days := [...]string{"Sat", "Sun"}  // len(days) == 2

[]T{x1, x2, ... xn}

tmp := [n]T{x1, x2, ... xn}
tmp[0 : n]

[...]Point{{1.5, -3.5}, {0, 0}}
[][]int{{1, 2, 3}, {4, 5}}
[][]Point{{{0, 1}, {1, 2}}}
map[string]Point{"orig": {0, 0}}
map[Point]string{{0, 0}: "orig"}

type PPoint *Point
[2]*Point{{1.5, -3.5}, {}}
[2]PPoint{{1.5, -3.5}, {}}

if x == (T{a, b, c}[i]) { ... }
if (x == T{a, b, c}[i]) { ... }

primes := []int{2, 3, 5, 7, 2147483647}

vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'y': true}

filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}

noteFrequency := map[string]float32{
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87,
}
