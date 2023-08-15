package main

s := make([]int, 1e3)
// s := make([]int, 1<<63)
// s := make([]int, 10, 0)
c := make(chan int, 10)
m := make(map[string]int, 100)

append(s S, x ...E) S

s0 := []int{0, 0}
s1 := append(s0, 2)  // s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)  // s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)  // s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)  // s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")  // t == []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...) // b == []byte{'b', 'a', 'r'}

copy(dst, src []T) int
copy(dst []byte, src string) int

var a [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])  // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])  // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")

delete(m, k)

complex(realPart, imaginaryPart floatT) complexT
real(complexT) floatT
imag(complexT) floatT

var a = complex(2, -2)
const b = complex(1.0, -1.4)
x := float32(math.Cos(math.Pi/2))
var c64 = complex(1, 0)
var s int = complex(1, 0)
// _ = complex(1, 2<<s)
var rl = real(c64)
var im = imag(a)
const c imag(b)
// _ = imag(3 << s)

func panic(interface{})
func recover() interface{}

panic(42)
panic("unreachable")
panic(Error("cannot parse"))

func protect(g func()) {
	defer func() {
		log.Println("done")

		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()

	log.Println("start")
	g()
}

print
println

SourceFile = PackageClause ";" { ImportDecl ";" } { TopLevelDecl ";" } .

PackageClause = "package" PackageName .
PackageName = identifier .

package math

ImportDecl = "import" ( ImportSpec | "(" { ImportSpec ";" } ")" ) .
ImportSpec = [ "." | PackageName ] ImportPath .
ImportPath = string_lit .

package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i  // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {  // Loop over values received from 'src'.
		if i % prime != 0 {
			dst <- i  // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int)  // Create a new channel.
	go generate(ch)       // STart generate() as a subprocess.

	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)

		ch = ch1
	}
}

func main() {
	sieve()
}
