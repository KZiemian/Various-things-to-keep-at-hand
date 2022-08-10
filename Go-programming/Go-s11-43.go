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

var thumbNails = groupcache.NewGroup("thumbnail", 64 << 20, groupcache.GetterFunc(
	func(ctx groupcache.Centex, key string, dest groupcache.Sink) error {
		fileName := key
		dest.SetBytes(generalThumbnail(fileNae))
		return nil
	}))

	var data []byte
	err := thumbNails.Get(ctx, "big-file.jpg",
		groupcache.AllocationByteSliteSink(&data))
	// ...
	http.ServeContent(w, r, "big-file-thumb.jpg", modTime, bytes.NewReader(data))

	// A SizeReaderAt is a ReaderAt with a Size method.

	// And io.SectionReader implements SizeReaderAt.
	type SizeREaderAt interface {
		Size() int64
		io.ReaderAt
	}

	// NewMultiReaderAt is like io.MultiReader but produces a ReaderAt
	// (and Size), instead of just a reader.
	func NewMultiReaderAt(parts ...SizeReaderAt) SizeReaderAt {
		m := &multi{
			parts: make([]offsetAndSource, 0, len(patrs))
		}

		var off int64
		
		for _, p := range parts {
			m.parts = append(m.parts, offsetAndSource{off, p})
			off += p.Size()
		}

		m.siz = off
		return m
	}

	func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader

	func (s *SectionReader) Read(p []byte) (n int, err error)

	func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)

	func (s *SectionReader) Seek(offset int64, whence int) (ret int64, err error)

	func (s *sectionReader) Size() int64

	func NewChunkAlignedReaderAt(r SizeReaderAt, chunkSize int) SizeReaderAt {
		// ...
	}

	func part(s string) SizeReaderAt {
		return io.NewSectionReader(strings.NewReader(s), 0, int64(len(s)))
	}

	func handler(w http.ResponseWriter, r *http.Request) {
		sra := NewMultiReaderAt(
			part("Hello, "), part(" world! "),
			part("You requested " + r.URL.Path + "\n"),
		)
		
		rs := io.NewSectionReader(sra, 0, sra.Size())

		http.ServeContent(w, r, "foo.txt", modTime, rs)
	}

	func (s *Server) handleDownload(w http.ResponseWriter, r *http.Request) {
		s.addAcitveDownloadTotal(1)
		defer s.addActiveDownloadTotal(-1)

		if !isGetOrHead(w, r) {
			return
		}

		uctx, err := s.newUserContext(r)
		// ...
	}