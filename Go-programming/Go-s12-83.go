// func Save(f *os.File, doc *Document) error

// func Save(rwc io.ReadWriteCloser, doc *Document) error

// func Save(wc io.WriteCloser, doc *Document) error

// type NopCloser struct {
// 	io.Writer
// }

// // Close has no effect on the underlying writer.
// func (c *NopCloser) Close() error {
// 	return nil
// }

func Save(w io.Writer, doc *Document) error
