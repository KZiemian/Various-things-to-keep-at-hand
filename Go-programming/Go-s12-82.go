// func main() {
// 	p := Person{
// 		Name: "Francesc",
// 		SSN:  "123456789",
// 	}

// 	b, err := json.MarshalIndent(p, "", "\t")

// 	if err != nil {
// 		log.Fatalf("Encode person: %v", err)
// 	}

// 	fmt.Printf("%s\n", b)

// 	var d Person

// 	if err := json.Unmarshal(b, &d); err != nil {
// 		log.Fatalf("Decode person: %v", err)
// 	}

// 	fmt.Println(d)
// }

// package painkiller

// type Pill int

// const (
// 	Placebo Pill = iota
// 	Aspirin
// 	Ibuprofen
// 	Paracetamol
// )

package main

import "fmt"

// type A struct {
// 	year int
// }

// func (a A) Greet() {
// 	fmt.Println("Hello GolangUK", a.year)
// }

// type B struct {
// 	A
// }

// func (b B) Greet() {
// 	fmt.Println("Welcome to GolangUK", b.year)
// }

// func main() {
// 	var a A

// 	a.year = 2016

// 	var b B

// 	b.year = 2016

// 	a.Greet()
// 	b.Greet()
// }

// type Cat struct {
// 	Name string
// }

// func (c Cat) Legs() int {
// 	return 4
// }

// func (c Cat) PrintLegs() {
// 	fmt.Printf("I have %d legs.\n", c.Legs())
// }

// type OctoCat struct {
// 	Cat
// }

// func (o OctoCat) Legs() int {
// 	return 5
// }

// func main() {
// 	var octo OctoCat

// 	fmt.Println(octo.Legs())
// 	octo.PrintLegs()
// }

// func PrintLegs(c Cat) {
// 	fmt.Printf("I have %d legs.\n", c.Legs())
// }

type Reader interface {
	// Read reads up to len(buf) bytes into buf.
	Read(buf []byte) (n int, err error)
}
