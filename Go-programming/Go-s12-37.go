type Complex interface {
	~complex64 | ~complex128
}

type Ordered interface {
	Integer | Float | ~string
}

// type Set[T any] map[T]struct{}

type Set[T comparable] map[T]struct{}

type Message interface {
	json.Marshaler
	json.Unmarshaler
}

func Call[Req, Resp Message](req Req) (resp Resp, error) {
	b, err := req.MarshalJSON()

	if err != nil {
		return resp, err
	}

	err := resp.UnmarshalJSON(b)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

type Request struct {
	// ...
}

func (m *Request) MarshalJSON() ([]byte, error) {
	// ...
}

func (m *Request) UnmarshalJSON(b []byte, error) {
	// ...
}

type Response struct {
	// ...
}

func (m *Response) MarshalJSON() ([]byte, error) {
	// ...
}

func (m *Response) UnmarshalJSON(b []byte, error) {
	// ...
}

resp, err := Call[Request, Response](req)

resp, err := Call[*Request, *Response](req)

package x

import (
	"encoding/json"
)

func Call[TODO](req Req) (resp Resp, error) {
	b, err := (&req).MarshalJSON()

	if err != nil {
		return resp, err
	}

	err := (&resp).UnmarshalJSON(b)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

type Message[T any] interface {
	*T
	json.Marshaler
	json.Unmarshaler
}

func Call[Req, Resp any, ReqP Message[Req], RespP Message[Resp]](req Req) (resp Resp, error) {
	b, err := ReqP(&req).MarshalJSON()

	if err != nil {
		return resp, err
	}

	err := RespP(&resp).UnmarshalJSON(b)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
