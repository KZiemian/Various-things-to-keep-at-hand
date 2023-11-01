package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	fileName := "words.txt"

// 	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY,
// 		0o644)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	if _, err := f.WriteString("cloud\n"); err != nil {
// 		log.Fatal(err)
// 	}
// }

// import (
// 	"io/ioutil"
// 	"log"
// )

// func main() {
// 	src := "words.txt"
// 	dest := "words_2.txt"

// 	bytesRead, err := ioutil.ReadFile(src)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = ioutil.WriteFile(dest, bytesRead, 0o644)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	var files []string

// 	root := "/home/janbodnar/Documents"

// 	err := filepath.Walk(root, func(path string, info os.FileInfo,
// 		err error) error {
// 		if err != nil {
// 			fmt.Println(err)

// 			return nil
// 		}

// 		if !info.IsDir() && filepath.Ext(path) == ".txt" {
// 			files = append(files, path)
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, file := range files {
// 		fmt.Println(file)
// 	}
// }

type Config struct {
	// ...
}

func WithReticulatedSplines(c *Config) {
	// ...
}

type Terrain struct {
	config Config
}

func NewTerrain(options ...func(*Config)) *Terrain {
	var t Terrain

	for _, option := range options {
		option(&t.config)
	}

	return &t
}

func main() {
	t := NewTerrain(WithReticulatedSplines)
	// ...
}
