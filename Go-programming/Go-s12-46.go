// Go Spec
x, y = f()

one, two, three = '一', '二', '三'

_ = x
x, _ = f()

a, b = b, a

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2

i = 0
x[i], i = 2, 1

x[0], x[0] = 1, 2

x[1], x[3] = 4, 5

type Point struct { x, y int }

var p *Point

x[2], p.x = 6, 7

i = 2
x = []int{3, 5, 7}

for i, x[i] = range x {
	break
}

IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ]

if x > max {
	x = max
}

if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}

SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .

ExprSwitchStmt = "switch"[ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
ExprCaseClause = ExprSwitchCase ":" StatementList .
ExprSwitchCase = "case" ExpressionList | "default" .

switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {
case x < 0:
	return -x
default: return x
}

switch {
case x < y: f1()

case x < z:
	f2()

case x == 4:
	f3()
}

switch x.(type) {
	// cases
}

TypeSwitchStmt = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClauses } "}" .
TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
TypeCaseClause = TypeSwitchCase ":" StatementList .
TypeSwitchCase = "case" TypeList | "default" .

switch i := x.(type) {
case nil:
	printString("x is nil")

case int:
	printInt(i)

case float64:
	printFloat64(i)

case func(int) float64:
	printFunction(i)

case bool, string:
	printString("type is bool or string")

default:
	printSTring("don't know the type")
}
