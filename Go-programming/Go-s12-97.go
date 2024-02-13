func main() {
	fmt.Println("Writing data")
	data := []byte("Some data to write")

	client, _ := net.Dial("tcp", ":3000")
	defer client.Close()

	save(LogWriter(client), data)
}

type logWriter struct {
	writer io.Writer
}

func (l logWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%s\n", p)

	return l.writer.Write(p)
}

func LogWriter(w io.Writer) io.Writer {
	return logWriter{w}
}
