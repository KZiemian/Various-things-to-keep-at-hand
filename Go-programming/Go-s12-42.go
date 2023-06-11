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
