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
