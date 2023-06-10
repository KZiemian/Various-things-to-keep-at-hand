t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)

(*T).Mp

func(tp *T, f float32) float32

(*T).Mv

func(tv *T, a int) int

t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)

(*T).Mp

func(tp *T, f float32) float32

(*T).Mv

func(tv *T, a int) int

type S struct {
	*T
}

type T int

func (t T) M() {
	print(t)
}

t := new(T)
s := S{T: t}
f := t.M
g := s.M
*t = 42

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
var pt *T
func makeT() T

t.Mv

func(int) int

t.Mv(7)
f := t.Mv; f(7)

pt.Mp

func(float32) float32

f := t.Mv; f(7)  // t.Mv(7)
f := pt.Mp; f(7) // pt.Mp(7)
f := pt.Mv; f(7) // (*pt).Mv(7)
f := t.Mp; f(7)  // (&t).Mp(7)
// f := makeT().Mp

var i interface { M(int) } = myVal
f := i.M; f(7) // i.M(7)

v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]

a[low : high]

a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]

a[2:]
a[:3]
a[:]

var a [10]int
s1 := a[3:7]
s2 := s1[1:4]
s2[1] = 42

var s []int
s3 := s[:0]
