package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// )

// func main() {
// 	content, err := ioutil.ReadFile("file.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("File content: %v.\n", content)
// 	fmt.Printf("Content type: %T.\n", content)
// 	fmt.Printf("File content as string: %vEnd of file.\n",
// 		string(content))
// }

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	f, err := os.Open("file.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)

// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// )

// func main() {
// 	fileName := "data.txt"

// 	val := "old\nfalcon\nsky\ncup\nforest\n"
// 	data := []byte(val)

// 	err := ioutil.WriteFile(fileName, data, 0o644)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("done")
// }

// import "fmt"

// func main() {
// 	fmt.Printf("0644: %v.\n", 0644)
// 	fmt.Printf("0o644: %v.\n", 0o644)
// }

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "data.txt"

	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	words := []string{"sky", "falcon", "rock", "hawk"}

	for _, word := range words {
		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}
