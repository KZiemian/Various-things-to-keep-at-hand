package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Printf("math.E: %.10f.\n", math.E)
	// fmt.Printf("math.Pi: %.10f.\n", math.Pi)
	// fmt.Printf("math.Phi: %.10f.\n", math.Phi)

	// fmt.Printf("math.Sqrt2: %.10f.\n", math.Sqrt2)
	// fmt.Printf("math.SqrtE: %.10f.\n", math.SqrtE)
	// fmt.Printf("math.SqrtPi: %.10f.\n", math.SqrtPi)
	// fmt.Printf("math.SqrtPhi: %.10f.\n", math.SqrtPhi)

	// fmt.Printf("math.Ln2: %.10f.\n", math.Ln2)
	// fmt.Printf("math.Log2E: %.10f.\n", math.Log2E)
	// fmt.Printf("1 / math.Ln2: %.10f.\n", 1 / math.Ln2)
	// fmt.Printf("math.Ln10: %.10f.\n", math.Ln10)
	// fmt.Printf("math.Log10E: %.10f.\n", math.Log10E)
	// fmt.Printf("1 / math.Ln10: %.10f.\n", 1 / math.Ln10)

	// fmt.Printf("math.MaxFloat32: %.10g.\n", math.MaxFloat32)
	// fmt.Printf("math.MaxFloat64: %.10g.\n", math.MaxFloat64)
}

v := x

if v == nil {
	i := v
	printString("x is nil")
} else, if i, isInt := v.(int); isInt {
	printInt(i)
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)

	if isBool || isString {
		i := v
		printString("type is bool or string")
	} else {
		i := v
		printString("don't know the type")
	}
}

func f[P any](x any) int {
	switch x.(type) {
	case P:
		return 0

	case string:
		return 1

	case []P:
		return 2

	case []byte:
		return 3

	default:
		return 4
	}
}

var v1 = f[string]("foo")
var v2 = f[byte]([]byte{})

ForStmt = "for" [ Condition | ForClause |RangeClause ] Block .
Condition = Expression .

for a < b {
	a *= 2
}

ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
InitStmt = SimpleStmt .
PostStmt = SimpleStmt .

for i := 0; i < 10; i++ {
	f(i)
}

for cond {
	S()
}
for ; cond ; {
	S()
}

for {
	S()
}

for true {
	S()
}
