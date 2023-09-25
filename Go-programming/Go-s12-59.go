package main

import "fmt"

// func main() {
// 	// a := make([]int, 2, 100)

// 	// fmt.Printf("a: %v.\n", a)
// 	// fmt.Printf("length: %v.\n", len(a))
// 	// fmt.Printf("capacity: %v.\n", cap(a))

// 	// a = append(a, 1)

// 	// fmt.Printf("a: %v.\n", a)
// 	// fmt.Printf("length: %v.\n", len(a))
// 	// fmt.Printf("capacity: %v.\n", cap(a))

// 	a := make([]int, 100)

// 	fmt.Printf("length: %v.\n", len(a))
// 	fmt.Printf("capacity: %v.\n", cap(a))

// 	a = append(a, 1)

// 	fmt.Printf("length: %v.\n", len(a))
// 	fmt.Printf("capacity: %v.\n", cap(a))
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5}

// 	idx := 1

// 	a = append(a[:idx], a[idx+1:]...)

// 	fmt.Println(a)
// }

// func main() {
// 	// mapVar := make(map[string]string, 10)

// 	// fmt.Printf("length: %v.\n", mapVar)

// 	favLangs := make(map[string]string)

// 	favLangs["Thomas"] = "golang"
// 	favLangs["Jane"] = "golang"
// 	favLangs["John"] = "Java"

// 	fmt.Printf("favLangs: %v.\n", favLangs)
// }

// func main() {
// 	favLangs := map[string]string{"Thomas": "golang", "Jane": "golang",
// 		"John": "Java"}

// 	fmt.Printf("favLangs: %v.\n", favLangs)

// 	fmt.Printf("favLangs[%v]: %v.\n", "Thomas", favLangs["Thomas"])
// 	fmt.Printf("favLangs[%v]: %v.\n", "thomas", favLangs["thomas"])
// }

// func main() {
// 	favLangs := map[string]string{"Thomas": "golang", "Jane": "golang",
// 		"John": "Java"}

// 	v, exists := favLangs["thomas"]

// 	fmt.Printf("v: %v.\n", v)
// 	fmt.Printf("v: %v.\n", exists)

// 	if exists {
// 		fmt.Println("a value is present")
// 	} else {
// 		fmt.Println("no value is present")
// 	}
// }

// func main() {
// 	favLangs := map[string]string{"Thomas": "golang", "Jane": "golang",
// 		"John": "Java"}

// 	fmt.Printf("favLangs: %v.\n", favLangs)

// 	delete(favLangs, "Jane")

// 	fmt.Printf("favLangs: %v.\n", favLangs)
// }
