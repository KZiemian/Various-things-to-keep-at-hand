resp, err := Call[Request, Response, *Request, *Response](req)

resp, err := Call[Request, Response](req)

func Sort[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func Example() {
	// f := Sort
	f := Sort[int]
}

type Set[A comparable] map[A]struct{}

// func (s Set[A]) Map[B comparable](f func(A) B) Set[B]

func Map[A, B comprable](s Set[A], f func(A) B) Set[B]

type X struct{}

func (X) F[T any](v T) {}

func FarAwayCode(x X) {
	fint := x.(interface{ F(int) })
}
