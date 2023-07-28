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
