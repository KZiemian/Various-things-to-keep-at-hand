pub.Publish[*Response](&Response{ /* ... */ })
// becomes
pub.Publish(&Response{ /* ... */ })

provider.Call[*dep.Request, *dep.Response](ctx, dep.ID, &dep.Request{ /* ... */ })

package provider

type ID[Request, Response any] string

type Provider[Request, Response any] interface {
	ID() ID[Request, Response]
	Run(context.Context, Publisher[Response], Request) error
	Close() error
}

type Publisher[Response any] interface {
	Publish(Response)
}

func Call[Req, Resp any](context.Context, ID[Req, Resp], Req) (Resp, error)

func Stream[Req, Resp any](context.Context, ID[Req, Resp], Req) Iterator[Resp]

package dep

const ID = provider.ID[*Request, *Response]("depProvider")

package foo

resp, err := provider.Call[*dep.Request, *Wrong](ctx, dep.ID, &dep.Request{ /* ... */ })

resp, err := provider.Call(ctx, dep.ID, &dep.Request{ /* ... */ })

type ID[Request, Response Message] string

type Provider[Request, Response Message] interface {
	ID() ID[Request, Response]
	Run(context.Context, Publisher[Response], Request) error
	Close() error
}

type Publisher[Response Message] interface {
	Publish(Response)
}

func Call[Req, Resp Message](context.Context, ID[Req, Resp], Req) (Resp, error)

func Stream[Req, Resp Message](context.Context, ID[Req, Resp], Req) Iterator[Resp]

// const ID = provider.ID[int, string]

func Call[Req, Resp Message](ctx context.Context, p ID[Req, Resp], r Req) (Resp, error) {
	buf, _ := r.MarshalJson()
	// r.Foo()
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type ByteSeq interface {
	~string | ~[]byte
}

func Concat[T ByteSeq](s ...T) T {
	var out []byte

	for _, piece := range s {
		out = append(out, pieace...)
	}

	return T(out)
}

Concat("foo", "bar", "baz")
Concat([]byte{102}, []byte{111}, []byte{111})
// Concat(42, 23, 1337)

package strings

func Join[S ~string](parts []S, sep S) S {
	p := []string(parts)
	joinde := strings.Join(p, string(sep))

	return S(joined)
}

type Path string

const Sep Path = "/"

func Join(parts ...Path) Path {
	return strings.Join(parts, Sep)
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~unit | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}
