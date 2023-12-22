func do() *doError {
	return nil
}

func main() {
	err := do()

	fmt.Println(err == nil)
}

func do() *doError {
	return nil
}

func wrapDo() error {
	return do()
}

func main() {
	err := wrapDo()

	fmt.Println(err == nil)
}

var ptrInt *int

ptrInt == nil

*ptrInt

type tree struct {
	v int
	l *tree
	r *tree
}

func (t *tree) Sum() int

func (t *tree) Sum() int {
	sum := t.v

	if t.l != nil {
		sum += t.l.Sum()
	}

	if t.r != nil {
		sum += t.r.Sum()
	}

	return sum
}

type person struct {}

func sayHi(p *person) {
	fmt.Println("hi")
}

func (p *person) sayHi() {
	fmt.Println("hi")
}

var p *person

p.sayHi() // hi

func (t *tree) Sum() int {
	if t == nil {
		return 0
	}

	return t.v + t.l.Sum() + t.r.Sum()
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}

	return fmt.Sprint(t.l, t.v, t.r)
}

func (t *tree) Find(v int) bool {
	if t == nil {
		return false
	}

	return t.v == v || t.l.Find(v) || t.r.Find(v)
}

var m map[t]u

len(m)

v, ok := m[i]

m[i] = x

func NewGet(url string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.Method, url, nil)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}
