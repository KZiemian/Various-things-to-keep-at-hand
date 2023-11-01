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

// func Save(w io.Writer, doc *Document) error

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	_, err := os.Stat("words.txt")

// 	if errors.Is(err, os.ErrNotExist) {
// 		fmt.Println("File does not exist.")
// 	} else {
// 		fmt.Println("File exists.")
// 	}

// 	fmt.Printf("err: %v.\n", err)
// 	fmt.Printf("err: %T.\n", err)
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.Create("empty.txt")

// 	defer file.Close()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("file created")
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	err := os.Remove("empty.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("file deleted")
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	fInfo, err := os.Stat("file.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fsize := fInfo.Size()

// 	fmt.Printf("The file size is %d bytes.\n", fsize)
// 	fmt.Printf("fsize: %T.\n", fsize)
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	fileName := "file.txt"

// 	fileInfo, err := os.Stat(fileName)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	mTime := fileInfo.ModTime()

// 	fmt.Printf("mTime: %v.\n", mTime)
// 	fmt.Printf("mTime type: %T.\n", mTime)
// }
