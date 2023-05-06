// const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer
// // and string constants
// const u, v float32 = 0, 3    // u = 0.0, v = 3.0

// const (
// 	Sunday = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Partyday
// 	numberOfDays   // this constant is not exported
// )

// const (
// 	c0 = iota  // c0 == 0
// 	c1 = iota  // c1 == 1
// 	c2 = iota  // c2 == 2
// )

// const (
// 	a = 1 << iota  // a == 1 (iota == 0)
// 	b = 1 << iota  // b == 2 (iota == 1)
// 	c = 3          // c == 3 (iota == 2, unused)
// 	d = 1 << iota  // d == 8 (iota == 3)
// )

// const (
// 	u          = iota * 42  // u == 0    (untyped integer constant)
// 	v  float64 = iota * 42  // v == 42.0 (float64 constant)
// 	w          = iota * 42  // w == 84   (untyped integer constant)
// )

// const x = iota  // x == 0
// const y = iota  // y == 0

// const (
// 	bit0, mask0 = 1 << iota, 1 << iota - 1  // bit0 == 1, mask0 == 0
// 	// (iota == 0)
// 	bit1, mask1                             // bit1 == 2, mask1 == 1
// 	// (iota == 1)
// 	_, _                                    // (iota == 2, unused)
// 	bit3, mask3                             // bit3 == 8, mask3 == 7
// 	// (iota = 3)
// )

// TypeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
// TypeSpec = AliasDecl | TypeDef .

// AliasDecl = identifier "=" Type .

// type (
// 	nodeList = []*Node // nodeList and []*Node are identical types
// 	Polar    = polar   // Polar and polar denote identical types
// )

// TypeDef = identifier [ TypeParameters ] Type .

// type (
// 	Point struct{ x, y float64 }  // Point and struct{ x, y float64 } are
// 	// different types
// 	polar Point                   // polar and Point denote different
// 	// types
// )

// type TreeNode struct {
// 	left, right *TreeNode
// 	value       any
// }

// type Block interface {
// 	BlockSize() int
// 	Encrypt(src, dst []byte)
// 	Decrypt(src, dst []byte)
// }

// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct {
	// Mutex fields
}

func (m *Mutex) Lock() {
	// Lock implementation
}

func (m *Mutex) Unlock() {
	// Unlock implementation
}

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of PtrMutex's underlying type *Mutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its embedded field Mutex.
type PrintableMutex struct {
	Mutex
}
