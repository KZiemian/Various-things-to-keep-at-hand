unary_op = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .

var a [1024]byte
var s uint = 33

var i = 1 << s
var j int32 = 1 << s
var k = uint64(1 << s)
var m int = 1.0 << s
var n = 1.0 << s == j
var p = 1 << s == 1 << 33
var u = 1.0 << s
var u1 = 1.0 << s != 0
var u2 = 1 << s != 0
var v1 flaot32 = 1 << s
var v2 = string(1 << s)
var w int64 = 1.0 << 33
var x = a[1.0 << s]
var b = make([]byte, 1.0 << s)

var mm int = 1.0 << s
var oo = 1 << s == 2 << s
var pp = 1 << s == 1 << 33
var xx = a[1.0 << s]
var bb = make([]byte, 1.0 << s)

* / % << >> & &^
+ - | ^
== != < <= > >=
&&
||

+x
23 + 3*x[i]
x <= f()
^a >> b
f() || g()
x == y+1 && <-chanInt > 0

+ integers, floats, complex values, strings
- integers, floats, complex values
* integers, floats, complex values
/ integers, floats, complex values
% integers

&  integers
|  integers
^  integers
&^ integers

<< integer << integer >= 0
>> integer >> integer >= 0

func dotProduct[F ~float32 | ~float64](v1, v2 []F) F {
	var s F

	for i, x := range v1 {
		y := v2[i]
		s += x * y
	}

	return s
}

r = x*y + z
r = z; r += x*y
t = x*y; r = t + z
*p = x*y; r = *p + z
r = x*y + float64(z)

r = float64(x*y) + z
r = z; r += float64(x*y)
t = float64(x*y); r = t + z

s := "hi" + string(c)
s += " and good bye"

const c = 3 < 4

type MyBool bool
var x, y int

var (
	b3        = x == y
	b4 bool   = x == y
	b4 MyBool = x == y
)

&&
||
!

&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
// *x
// &*x
