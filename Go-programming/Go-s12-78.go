// package main

// import "fmt"

// func main() {
// 	i := 1
// 	var ptr1 *int = &i
// 	var ptr2 **int = &ptr1

// 	fmt.Printf("i: %T, %v.\n", i, i)
// 	fmt.Printf("ptr1: %T, %v.\n", ptr1, ptr1)
// 	fmt.Printf("ptr2: %T, %v.\n", ptr2, ptr2)
// }

package main

import (
	"fmt"
	"errors"
)

// common HTTP status codes
var NotFoundHTTPCode = errors.New("404")
var UnauthorizedHTTPCode = errors.New("401")

// database errors
var RecordNotFoundErr = errors.New("DB: record not found")
var AffectedRecordsMismatchErr = errors.New("DB: affected records mismatch")

// HTTP client errors
var ResourceNotFoundErr = errors.New("HTTP client: resource not found")
var ResourceUnauthorizedErr = errors.New("HTTP client: unauthorized")

// application errors (the new features)
var UserNotFoundErr = fmt.Errorf("user not found: %w (%w)",
	RecordNotFoundErr, NotFoundHTTPCode)
var OtherResourceUnauthorizedErr = fmt.Errorf("unauthorized call: %w (%w)",
	ResourceUnauthorizedErr, UnauthorizedHTTPCode)

func handleError(err error) {
	if errors.Is(err, NotFoundHTTPCode) {
		fmt.Println("Will return 404")
	} else if errors.Is(err, UnauthorizedHTTPCode) {
		fmt.Println("Will return 401")
	} else {
		fmt.Println("Will return 500")
	}

	fmt.Println(err.Error())
}

func main() {
	handleError(UserNotFoundErr)
	handleError(OtherResourceUnauthorizedErr)
}

{
	"name": "Gopher",
	"birthdate": "2009/11/10",
	"shirt-size": "XS",
}

type Person struct {
	Name string
	Born time.Time
	Size ShirtSize
}

type ShirtSize byte

const (
	NA ShirtsSize = iota
	XS
	S
	M
	L
	XL
)

func (p *Person) Parse(s string) error {
	fields := map[string]string{}

	dec := json.NewDecoder(strings.NewReader(s))

	if err := dec.Decode(&fields); err != nil {
		return fmt.Errof("decode person: %w", err)
	}

	// p.Name = fields["name"]
}

func ParseShirtSize(s string) (ShirtSize, error) {
	sizes := map[string]ShirtSize{"XS", XS, "S": S, "M": M, "L": L,
		"XL": XL}

	giveSize, ok := sizes[s]

	if !ok {
		return NA, fmt.Errorf("invalid ShirtSize %q", s)
	}

	return giveSize, nil
}
