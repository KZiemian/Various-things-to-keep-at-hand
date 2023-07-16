v1 := <-ch
v2 = <-ch
f(<-ch)

x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch

Conversion = Type "(" Expression [ "," ] ")" .

*Point(p) // *(Point(p))
(*Point)(p)
<-chan int(c) // <-(chan int(c))
(<-chan int)(c) // <-chan int
func()(x)
(func())(x)
(func() int)(x)
func() int(x)

uint(iota)
float32(2.7182818284)
complex128(1)
float32(0.49999999)
float64(-1e-1000)
string('x')
string(0x266c)
myString("foo" + "bar")
string([]byte{'a'})
(*int)(nil)

func f[P ~float32 | ~float64]() {
	... P(1.1) ...
}

type Person struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

var data *struct {
	Name    string  `json:"name"`
	Address *struct {
		Street string `json:"street"`
		City   string `json:"city"`
	} `json:"adress"`
}

var person = (*Person)(data)

string([]byte{})
string([]byte(nil))

// []byte("hell")  // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
// []byte("") // []byte{}

bytes("hell")  // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}

[]myByte("world!") // []myByte{'w', 'o', 'r', 'l', 'd', '!'}
[]myByte(myString("???")) // []myByte{'\xf0', '\x9f', '\x8c', '\x8f'}

const a = 2 + 3.0  // a == 5.0
const b = 15 / 4   // b == 3
const c = 15 / 4.0 // c == 3.75
const theta float64 = 3 / 2  // theta == 1.0
const Pi float64 = 3 / 2. // Pi == 1.5
const d = 1 << 3.0 // d == 8
const e = 1.0 << 3 // e == 8
// const f = int32(1) << 33 // illega;
// const g = float64(2) >> 1
const h = "foo" > "bar" // h == true
const j = true
const k = 'w' + i // k = 'x'
const l = "hi"
const m = string(k)  // m == "x"
const Sigma = 1 - 0.707i
const Delta = Sigma + 2.0e-4
const Phi = iota*1i - 1/1i

const ic = complex(0, c) // ic == 3.75i
const i0 = complex(0, theta)

const Huge = 1 << 100
const Four int3 = Huge >> 98

// uint(-1)
// int(3.14)
// int64(Huge)
// Four * 300
// Four * 100

^1
// uint8(^1)
^uint8(1) // 0xFF ^ uint8(1) = uint(0xFE)
int8(^1) // int8(-2)
^int8(1) // -1 ^ int8(1) = -2

Statement =
	Declaration | LabeledStmt | SimpleStmt |
	GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
	FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt |
	ForStmt | DeferStmt .
