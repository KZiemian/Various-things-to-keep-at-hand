// func First(query string, replicas ...Search) Result {
// 	c := make(chan Result)

// 	searchReplica := func(i int) {
// 		c <- replicas[i](query)
// 	}

// 	for i := range replicas {
// 		go searchReplica(i)
// 	}

// 	return <-c
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	start := time.Now()

// 	result := First("golang",
// 		fakeSearch("replica 1"),
// 		fakeSearch("replica 2"))

// 	elapsed := time.Since(start)

// 	fmt.Println(result)
// 	fmt.Println(elapsed)
// }

// c := make(chan Result)

// go func() {
// 	c <- First(query, Web1, Web2)
// }()

// go func() {
// 	c <- First(query, Image1, Image2)
// }()

// go func() {
// 	c <- First(query, Video1, Video2)
// }

// timeout := time.After(80 * time.Millisecond)

// for i := 0; i < 3; i++ {
// 	select {
// 	case result := <-c:
// 		results = append(results, result)

// 	case <-timeout:
// 		fmt.Println("timed out")

// 		return
// 	}
// }

// return

// int_lit = decimal_lit | binary_lit | octal_lit | hex_lit .
// decimal_lit = "0" | ( "1" ... "9" ) [ [ "_" ] decimal_digits ] .
// binary_lit = "0" ( "b" | "B" ) [ "_" ] binary_digits .
// octal_lit = "0" [ "o" | "O" ] [ "_" ] octal_digits .
// hex_lit = "0" ( "x" | "X" ) [ "_" ] hex_digits .

// decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
// binary_digits = binary_digit { [ "_" ] binary_digit } .
// octal_digits = octal_digit { [ "_" ] octal_digit } .
// hex_digits = hex_digit { [ "_" ] hex_digit } .

// 42
// 4_2

package main

import "fmt"

func main() {
	// fmt.Printf("42: %v.\n", 42)
	// fmt.Printf("4_2: %v.\n", 4_2)
	// fmt.Printf("0600: %v.\n", 0600)
	// fmt.Printf("0_600: %v.\n", 0_600)
	// fmt.Printf("0o600: %v.\n", 0o600)
	// fmt.Printf("0O600: %v.\n", 0O600)
	// fmt.Printf("0xBadFace: %v.\n", 0xBadFace)
	// fmt.Printf("0xBad_Face: %v.\n", 0xBad_Face)
	// fmt.Printf("0x_67_7a_2f_cc_40_c6: %v.\n", 0x67_7a_2f_cc_40_c6)

	// fmt.Printf("0.: %v.\n", 0.0)
	// fmt.Printf("72.40: %v.\n", 72.40)
	// fmt.Printf("072.40: %v.\n", 072.40)
	// fmt.Printf("2.71828: %v.\n", 2.71828)
	// fmt.Printf("1.e+0: %v.\n", 1.e+0)
	// fmt.Printf("1.e0: %v.\n", 1.e0)
	// fmt.Printf("6.67428e-11: %v.\n", 6.67428e-11)
	// fmt.Printf("1E6: %v.\n", 1E6)
	// fmt.Printf(".25: %v.\n", .25)
	// fmt.Printf(".12345E+5: %v.\n", .12345E+5)
	// fmt.Printf("1_5.: %v.\n", 1_5.)
	// fmt.Printf("0.15e+0_2: %v.\n", 0.15e+0_2)

	fmt.Printf("0x1p-2: %v.\n", 0x1p-2)
	fmt.Printf("0x2.p10: %v.\n", 0x2.p10)
	fmt.Printf("0x1.Fp+0: %v.\n", 0x1.Fp+0)
	fmt.Printf("15.0 / 16.0: %v.\n", 15.0 / 16.0)
	fmt.Printf("0X.8p-0: %v.\n", 0X.8p-0)
	fmt.Printf("0X.8p0: %v.\n", 0X.8p0)
	fmt.Printf("0X_1FFFP-16: %v.\n", 0X1FFFP-16)
	fmt.Printf("0x15e-2: %v.\n", 0x15e-2)
	fmt.Printf("0x15e: %v.\n", 0x15e)
}
