// func (m *Mux) SendMsg(msg string) {
// 	result := make(chan error, 1)

// 	m.ops <- func(m map[net.Addr]net.Conn) {
// 		for _, conn := range m {
// 			err := io.WriteString(conn, msg)

// 			if err != nil {
// 				result <- err

// 				return
// 			}
// 		}
// 	}

// 	return <-result
// }

// func (m *Mux) loop() {
// 	conns := make(map[net.Addr]net.Conn)

// 	for _, op := range m.ops {
// 		op(conns)
// 	}
// }

// func (m *Mux) PrivateMsg(add net.Addr, msg string) error {
// 	result := make(chan net.Conn, 1)

// 	m.ops <- func(m map[net.Addr]net.Conn) {
// 		result <- m[addr]
// 	}

// 	conn := <-result

// 	if conn == nil {
// 		return errors.Errorf(
// 			"client %v not registered", addr)
// 	}

// 	return io.WriteString(conn, msg)
// }

// package main

// import "fmt"

// func send(message string) {
// 	fmt.Println(message)
// }

// func main() {
// 	send("Hi, it's me.")

// 	func (message string) {
// 		fmt.Println(message)
// 	}("Hi, it's me!")
// }

// func main() {
// 	sum := func(a, b, c int) int {
// 		return a + b + c
// 	}(3, 5, 7)

// 	fmt.Printf("3 + 5 + 7 = %v.\n", sum)
// }

// func intSeq() func() int {
// 	i := 0

// 	return func() int {
// 		i++
// 		return i
// 	}
// }

// func main() {
// 	nextInt := intSeq()

// 	fmt.Printf("nextInt(), 1: %v.\n", nextInt())
// 	fmt.Printf("nextInt(), 2: %v.\n", nextInt())
// 	fmt.Printf("nextInt(), 3: %v.\n", nextInt())
// 	fmt.Printf("nextInt(), 4: %v.\n", nextInt())

// 	nextInt2 := intSeq()

// 	fmt.Printf("nextInt2(): %v.\n", nextInt2())
// }

// package main

// import "fmt"

// func fibonacci() func() int {
// 	a := 0
// 	b := 1

// 	return func() int {
// 		a, b = b, a + b

// 		return b - a
// 	}
// }

// func main() {
// 	f := fibonacci()

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }
