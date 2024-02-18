// resp, err := Call[Request, Response, *Request, *Response](req)

// resp, err := Call[Request, Response](req)

// func Sort[T constraints.Ordered](s []T) {
// 	sort.Slice(s, func(i, j int) bool {
// 		return s[i] < s[j]
// 	})
// }

// func Example() {
// 	// f := Sort
// 	f := Sort[int]
// }

// type Set[A comparable] map[A]struct{}

// // func (s Set[A]) Map[B comparable](f func(A) B) Set[B]

// func Map[A, B comprable](s Set[A], f func(A) B) Set[B]

// type X struct{}

// func (X) F[T any](v T) {}

// func FarAwayCode(x X) {
// 	fint := x.(interface{ F(int) })
// }

// func save(f *os.File, data []byte) error {
// 	_, err := f.Write(data)

// 	return err
// }

// func Test_Save(t *testing.T) {
// 	testdata := "Test data"
// 	f, err := ioutil.TempFile("/tmp", "data")

// 	if err != nil {
// 		t.Fatal("Error creating test file.")
// 	}
// 	defer os.Remove(f.Name())

// 	if err := save(f, []byte(testdata)); err != nil {
// 		t.Fatal("Error saving data")
// 	}
// 	f.Close()

// 	data, _ := ioutil.ReadFile(f.Name())

// 	if string(data) != testdata {
// 		t.Errorf("Expected %s, got %s.\n", testdata, string(data))
// 	}
// }

// func save(w io.Writer, data []byte) error {
// 	_, err := w.Write(data)

// 	return err
// }

// func Test_Save(t *testing.T) {
// 	buf := &bytes.Buffer{}
// 	testdata := "Test data"

// 	if err := save(buf, []byte(testdata)); err != nil {
// 		t.Fatal("Error savign data.")
// 	}

// 	if buf.String() != testdata {
// 		t.Errorf("Expected %s, got %s.\n", testdata, buf.String())
// 	}
// }

// package main

// import "fmt"

// const (
// 	hApprox float64 = 6.626_070_15
// 	cApprox float64 = 2.997_924_58
// 	piApprox float64 = 3.141_592_653_5
// )

// func main() {
// 	casimirForce := piApprox * hApprox * cApprox / 480.0

// 	fmt.Printf("F = %v 10^-6.\n", casimirForce)
// }

package main

import "fmt"

const (
	hApprox float64 = 6.626_0
	cApprox float64 = 2.997_9
	piApprox float64 = 3.1415
)

func main() {
	casimirForce := piApprox * hApprox * cApprox / 480.0

	fmt.Printf("F = %v 10^-7.\n", 10 * casimirForce)
}
