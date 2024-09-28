// package main

// import "fmt"

// type Jedi struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// 	Rank      string
// }

// func (j Jedi) String() string {
// 	return j.FirstName + " " + j.LastName
// }

// func main() {
// 	luke := Jedi{
// 		FirstName: "Luke",
// 		LastName:  "Skywalker",
// 		Age:       22,
// 		Rank:      "Master",
// 	}

// 	fmt.Printf("luke: %v.\n", luke)
// }

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// func main(){
// 	fmt.Println("Writing data.")

// 	data := []byte("Some data to write.")

// 	f, _ := ioutil.TempFile("./", "data.txt")
// 	defer f.Close()

// 	save(f, data)
// }

// func save(f *os.File, data []byte) error {
// 	_, err := f.Write(data)

// 	return err
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// )

// func main() {
// 	fmt.Println("Writing data")

// 	data := []byte("Some data to write.")

// 	f, _ := ioutil.TempFile("./", "data.txt")

// 	defer f.Close()

// 	save(f, data)
// }

// func save(w io.Writer, data []byte) error {
// 	_, err := w.Write(data)

// 	return err
// }

// func main() {
// 	ln, err := net.Listen("tcp", ":3000")

// 	if err != nil {
// 		os.Exit()
// 	}

// 	fmt.Println("Listening on port 3000.")

// 	for {
// 		conn, err := ln.Accept()

// 		if err != nil {
// 			os.Exit()
// 		}

// 		go handleConnection(conn)
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	msg, err := ioutil.ReadAll(conn)

// 	if err != nil {
// 		os.Exit(1)
// 	}

// 	fmt.Printf("%s\n", msg)
// 	conn.Close()
// }

package main

import (
	"fmt"
	// "math"
)

func main() {
	eApprox := 2.7182

	fmt.Printf("eApprox^2 = %v.\n", eApprox * eApprox)
}
