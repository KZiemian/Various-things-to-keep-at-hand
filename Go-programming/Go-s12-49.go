package main

func complexF3() (re float64, im float64) {
	re = 7.0
	im = 4.0

	return
}

func (devnull) Write(p []byte) (n int, _ error) {
	n = len(p)

	return
}

func f(n int) (res int, err error) {
	if _, err := f(n - 1); err != nil {
		return // invalid return statement: err is shadowed
	}

	return
}

BreakStmt = "break" [ Label ] .

OuterLoop:
for i = 0; i < n; i++ {
	for j = 0; j < m; j++ {
		switch a[i][j] {
		case nil:
			state = Error
			break OuterLoop

		case item:
			state = Found
			break OuterLoop
		}
	}
}

RowLoop:
for y, row := range rows {
	for x, data := range row {
		if data == endOfRow {
			continue RowLoop
		}

		row[x] = data + bias(x, y)
	}
}

GotoStmt = "goto" Label .

goto Error

FallthroughtStmt = "fallthrough"

DeferStmt = "defer" Expression .

lock(l)

defer unlock(l)


for i := 0; i <= 3; i++ {
	defer fmt.Print(i)
}

func f() (result int) {
	defer func() {
		result *= 7
	}()

	return 6
}

const (
	c1 = imag(2i)
	c2 = len([10]float64{2})
	c3 = len([10]float64{c1})
	c4 = len([10]float64{imag(2i)})
	// c5 = len([10]float64{imag(z)})
)

var z complex128

new(T)

type S struct { a int; b float64 }
new(S)

make(T, n)
make(T, n, m)

make(T)
make(T, n)

make(T)
make(T, n)

s := make([]int, 10, 100)
