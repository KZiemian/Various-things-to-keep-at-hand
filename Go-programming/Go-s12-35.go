package provider

type ID string

type Provider interface {
	ID() ID
	Run(ctx context.Context, pub Publisher, req Message) error
	RequestType() reflect.Type
	ResponseType() reflect.Type
	Close() error
}

type Publisher interface{ Publish(Message) }

func Call(ctx context.Context, p ID, req Message) (Message, error)

func Stream(ctx context.Context, p ID, req Message) ResponseIterator

package foo

const ID = provider.ID("fooProvider")

type Request struct {
	// ...
}

type Response struct {
	// ...
}

func New() provider {
	return &prov{}
}

type prov struct {
	// ...
}

func (p *prov) ID() provider.ID {
	return ID
}

func (p *prov) RequestType() reflect.Type {
	return reflect.TypeOf(Request)
}

func (p *prov) ResponseType() reflect.Type {
	return reflect.TypeOf(Response)
}

func (p *prov) Run(ctx context.Context, pub provider.Publisher, req Message) error {
	r := req.(*Request)

	msg, err := provider.Call(ctx, dep.ID, &dep.Request{ /* ... */ })

	if err != nil {
		return nil
	}

	subresp := msg.(*dep.Response)

	pub.Publish(&Response{ /* ... */ })

	return nil
}

package provider

type Pulbisher[Response any] interface {
	Pulish(Response)
}

type Provider[Request, Response any] interface {
	ID() ID
	Run(context.Context, Publisher[Response], Requeste) error
	Close() error
}

func Call[Request, Response any](context.Context, ID, Request) (Response, error)

func Stream[Request, Reponse any](context.Context, ID, Request) Iterator[Response]

type Iterator[Response any] interface {
	Next() bool
	Resp() Response
}

package foo

const ID = provider.ID("fooProvider")

type Request struct { /* ... */ }
type Response struct { /* ... */ }

func New() provider.Provider[*Request, *Response] {
	return &prov{}
}

func (p *prov) ID() provider.ID {
	return ID
}

func (p *prov) Run(ctx context.Context, pub provider.Publisher[*Response],
	req *Request) error {
	resp, err := provider.Call[*dep.Request, *dep.Response](ctx, dep.ID, &dep.Request{ /* ... */ })

	if err != nil {
		return err
	}

	pub.Publish[*Response](&Response{ /* ... */ })

	return nil
}
