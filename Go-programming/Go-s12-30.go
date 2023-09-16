// // MyBlock is an interface type that has the same method set as Block.
// type MyBlock Block

// type TimeZone int

// const (
// 	EST TimeZone = -(5 + iota)
// 	CST
// 	MST
// 	PST
// )

// func (tz TimeZone) String() string {
// 	return fmt.Sprintf("GMT%+dh", tz)
// }

// type List[T any] struct {
// 	next  *List[T]
// 	value T
// }

// type T[P any] P   // illegal: P is a type parameter

// func f[T any]() {
// 	type L T  // illegal: T is a type parameter declared by the enclosing
// 	// function
// }

// // The method Len returns the number of elements in the linked list l.
// func (l *List[T]) Len() int {
// 	// ...
// }

// TypeParameters = "[" TypeParamList [ "," ] "]" .
// TypeParamList  = TypeParamDecl { "," TypeParamDecl } .
// TypeParamDecl  = IdentifierList TypeConstraint .

// [P any]
// [S interface{ ~[]byte | string }]
// [S ~[]E, E any]
// [P Constraint[int]]
// [_ any]

type T[P *C] ...
type T[P (C)] ...
type T[P *C|Q] ...

type T[P interface{*C}] ...
type T[P *C,] ...

type T1[P T1[P]] ...                   // illegal: T1 refers to itself
type T2[P interface{ T2[int] }] ...    // illegal: T2 refers to itself
type T3[P interface{ m(T3[int]) }] ... // illegal: T3 refers to itself
type T4[P T5[P]] ...                   // illegal: T4 refers to T5 and
type T5[P T4[P]] ...                   //          T5 refers to T4

type T6[P int] struct{ f *T6[P] }      // ok: reference to T6 is not in type
// parameter list

TypeDef = identifier [ TypeParmeters ] Type .

type (
	Point struct{ x, y float64 }
	polar Point
)

type TreeNode struct{
	left, right *TreeNode
	value any
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

type Mutex struct { /* Mutex fields */ }

func (m *Mutex) Lock() { /* Lock implementation */ }
func (m *Mutex) Unlock() { /* Unlock implementation */ }

type NewMutex Mutex

type PtrMutex *Mutex

type PrintableMutex struct {
	Mutex
}

type MyBlock Block

type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)
