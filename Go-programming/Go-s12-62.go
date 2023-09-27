package main

import "fmt"

// func main() {
// 	var a int = 42
// 	var b *int = &a

// 	fmt.Printf("&a: %v. b: %v.\n", &a, b)
// 	fmt.Printf("a: %v. *b: %v.\n", a, *b)
// }

// func main() {
// 	var a int = 42
// 	var b *int = &a

// 	fmt.Printf("a: %v. *b: %v.\n", a, *b)

// 	a = 41

// 	fmt.Printf("a: %v. *b: %v.\n", a, *b)

// 	*b = 37

// 	fmt.Printf("a: %v. *b: %v.\n", a, *b)
// }

// func main() {
// 	var a *int

// 	fmt.Printf("a: %v.\n", a)
// 	// fmt.Printf("*a: %v.\n", *a)
// }

type Person struct {
	Name     string
	Lastname string
	Age      int
	Address  string
}

type Employee struct {
	Person Person
	Salary int
}

type EmployeeEmbbeded struct {
	Person
	Salary int
}

// func main() {
// 	p := Person{
// 		Name:     "Thomas",
// 		Lastname: "Langhorst",
// 		Age:      34,
// 		Address:  "127.0.0.1",
// 	}

// 	fmt.Printf("p: %+v.\n", p)
// 	fmt.Printf("p.Name: %v.\n", p.Name)
// 	fmt.Printf("p.Name: %+v.\n", p.Name)
// }

// func main() {
// 	p1 := Person{
// 		Name:     "Thomas",
// 		Lastname: "Langhorst",
// 		Age:      34,
// 		Address:  "127.0.0.1",
// 	}

// 	p2 := p1
// 	p2.Name = "Definietly not"

// 	fmt.Printf("p1: %+v.\n", p1)
// 	fmt.Printf("p2: %+v.\n", p2)
// }

// func main() {
// 	p1 := Person{
// 		Name:     "Thomas",
// 		Lastname: "Langhorst",
// 		Age:      34,
// 		Address:  "127.0.0.1",
// 	}

// 	ptrPerson := &p1

// 	ptrPerson.Name = "Defienietly not"

// 	fmt.Printf("%+v.\n", p1)
// 	fmt.Printf("%+v.\n", ptrPerson)
// }

func main() {
	p := Person{
		Name:     "Thomas",
		Lastname: "Langhorst",
		Age:      34,
		Address:  "127.0.0.1",
	}

	e := Employee{
		Person: p,
		Salary: 999,
	}

	eEmbbeded := EmployeeEmbbeded{
		Person: p,
		Salary: 999,
	}

	// fmt.Printf("Full name of person p: %v.\n", p.FullName())
	// fmt.Printf("Full name of person p: %v.\n", FullName(p))

	fmt.Printf("p: %+v.\n", p)
	fmt.Printf("e: %+v.\n", e)

	fmt.Printf("e.Person.Fullname(): %v.\n", e.Person.FullName())

	fmt.Printf("eEmbbeded.FullName(): %v.\n", eEmbbeded.FullName())
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.Lastname)
}

func FullName(p Person) string {
	return fmt.Sprintf("%s %s", p.Name, p.Lastname)
}
