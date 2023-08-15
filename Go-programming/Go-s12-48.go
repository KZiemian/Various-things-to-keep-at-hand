package main

RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .

var testdata *struct {
	a *[7]int
}

for i, _ := range testdata.a {
	f(i)
}

var a [10]string

for i, x := range s {
	g(i, s)
}

var key string
var val interface{}
m := map[string]int{"mon": 0, "tue": 1, "wed": 2, "thu": 3, "fri": 4,
	"sat": 5, "sun": 6}

for key, val = range m {
	h(key, val)
}

var ch chan Work = producer()

for w := range ch {
	doWork(w)
}

for range ch {}

GoStmt = "go" Expression .

go Server()
go func(ch chan<- bool) {
	for {
		sleep(10);
		ch <- true
	}
}(c)

SelectStmt = "select" "{" { CommClause } "}" .
CommClause = CommCase ":" StatementList .
CommCase = "case" ( SendStmt | RecvStmt ) | "default" .
RecvStmt = [ ExpressionList "=" | IdentifierList ":=" ] RecvExpr .
RecvExpr = Expression .

var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int

select {
case i1 = <-c1:
	print("received ", i1, " from c1\n")

case c2 <- i2:
	print("sent ", i2, " to c2\n")

case i3, ok := (<-c3):
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}

case a[f()] = <-c4:

default:
	print("no communication\n")
}

for {
	select {
	case c <- 0:

	case c <-1:
	}
}

select {}

ReturnStmt = "return" [ ExpressionList ] .

func noResult() {
	return
}

func simpleF() int {
	return 2
}

func complexF1() (re float64, im float64) {
	return -7.0, -4.0
}

func complexF2() (re float64, im float64) {
	return complexF1()
}
