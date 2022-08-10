func ServerContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)

type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Seeker interface {
	Seek(offset int64, whence int) (ret int64, err error)
}

me := "http://10.0.0.1"
peers := groupcache.NewHTTPPool(me)

// Whenever peers change:
peers.Set("http://10.0.0.1", "http://10.0.0.2", "http://10.0.0.3")