// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go say("world")

// 	say("hello")
// }


// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)

// 		fmt.Println(s)
// 	}
// }

// FunctionType  = "func" Signature .
// Signature     = Parameters [ Result ] .
// Result        = Parameters | Type .
// Parameters    = "(" [ ParameterList [ "," ] ] ")" .
// ParameterList = ParameterDecl { "," ParameterDecl } .
// ParameterDecl = [ IdentifierList ] [ "... "] Type .

// func()
// func(x int) int
// func(a, _ int, z float32) bool
// func(a, b int, z float32) (bool)
// func(prefix string, values ...int)
// func(a, b int, z float64, opt ...inteface{}) (success bool)
// func(int, int, float64) (float64, *[]int)
// func(n int) func(p *T)

// InterfaceType  = "interface" "{" { InterfaceElem ";" } "}" .
// InterfaceElem  = MethodElem | TypeElem .
// MethodElem     = MethodName Signature .
// MethodNAme     = identifier .
// TypeElem       = TypeTerm { "|" TypeTerm } .
// TypeTerm       = Type | UnderlyingType .
// UnderlyingType = "~" Type .

// // A simple File interface.
// interface {
// 	Read([]byte) (int, error)
// 	Write([]byte) (int, error)
// 	Close() error
// }

// interface {
// 	String() string
// 	String() string  // illegal: String not unique
// 	_(x int)         // illegal: method must have non-blank name
// }

// func (p T) Read(p []byte) (n int, err error)
// func (p T) Write(p []byte) (n int, err error)
// func (p T) Close() error

// interface{}

// type Locker interface {
// 	Lock()
// 	Unlock()
// }

// func (p T) Lock() {
// 	// ...
// }

// func (p T) Unlock() {
// 	// ...
// }

type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Write interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
	Reader // includes methods of Reader in ReadWriter's method set
	Writer // includes methods of Writer in ReadWriter's method set
}

type ReadCloser interface {
	Reader  // includes methds of Reader in ReadCloser's method set
	Close() // illegal: signatures of Reader.Close and Close are different
}

// An inteface representing only the type int.
interface {
	int
}
