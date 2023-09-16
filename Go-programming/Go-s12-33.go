FunctionLit = "func" Signature FunctionBody .

func(a, b int, z float64) bool {
	return a*b < int(z)
}

f := func(x, y int) int {
	return x + y
}

func(ch chan int) { ch <- ACK }(replyChan)

PrimaryExpr =
	Operand |
		Conversion |
		MethodExpr |
		PrimaryExpr Selector |
		PrimaryExpr Index |
		PrimaryExpr Slice |
		PrimaryExpr TypeAssertion |
		PrimaryExpr Argumetns .

Selector = "." identifier .
Index = "[" Expression [ "," ] "]" .
Slice = "[" [ Expression ] ":" [ Expression ] "]" |
"[" [ Expression ] ":" Expression ":" Expression "]" .
TypeAssertion = "." "(" Type ")" .
Arguments = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ]
[ "," ] ] ")" .

x
2
(s + ".txt")
f(3.1415, true)
Point{1, 2}
m["foo"]
s[i : j + 1]
obj.color
f.p[i].x()

type T0 struct {
	x int
}

func (*T0) M0()

type T1 struct {
	y int
}

func (T1) M1()

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2()

type Q *T2

var t T2
var p *T2
var q Q = p

t.z  // t.z
t.y  // t.T1.y
t.x  // (*t.T0).x

p.z  // (*p).z
p.y  // (*p).T1.y
p.x  // (*(*p).T0).x

q.x  // (*(*q).T0).x

p.M0()  // ((*p).T0).M0()
p.M1()  // ((*p).T1).M1()
p.M2()  // p.M2()
t.M2()  // (&t).M2()

MethodExpr = ReceiverType "." MethodName .
ReceiverType = Type .

type T struct {
	a int
}

func (tv T) Mv(a int) int {
	return 0
}

func (tp *T) Mp(f float32) float32 {
	return 1
}

var t T

T.Mv

func(tv T, a int) int
