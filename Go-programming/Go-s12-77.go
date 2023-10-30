// func Query(conns []Conn, query string) Result {
// 	ch := make(chan Result, len(conns))  // buffered

// 	for _, conn := range conns {
// 		go func(c Conn) {
// 			ch <- c.DoQuery(qurey):
// 		}(conn)
// 	}

// 	return <-ch
// }

type error interface {
	Error() string
}

type err struct {
	s string
}

func (e *err) Error() string {
	return e.s
}

if err == io.EOR {
	// ...
}

if nerr, ok := err.(net.Error); ok {
	// ...
}

type WrappedError interface {
	Unwrap() error
}

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

var RecordNotFoundErr = errors.New("not found")
const name, id = "lzap", 13

werr := fmt.Errorf("unknown user %q (id %d): %w", name, id,
	recordNotFoundErr)

fmt.Println(werr.Error())

if errors.Is(err, RecordNotFoundErr) {
	// ...
}

err1 := errors.New("err1")
err2 := errors.New("err2")

err := errors.Join(err1, err2)

fmt.Println(err)

err1 := errors.New("err1")
err2 := errors.New("err2")

err := fmt.Errorf("%w + %w", err1, err2)

fmt.Println(err)

type joinError struct {
	errs []error
}

func (e *joinError) Error() string {
	// concatenate errors with a new line character
}

func (e *joinError) Unwrap() []error {
	return e.errs
}

type MultiWrappedError interface {
	Unwrap() []error
}

err1 := errors.New("err1")
err2 := errors.New("err2")

err := errors.Join(err1, err2)
unwrapped := errors.Unwrap(err)

fmt.Println(unwrapped)
