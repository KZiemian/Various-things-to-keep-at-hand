// func (m *Mux) Add(conn net.Conn) {
// 	m.mu.Lock()

// 	defer m.mu.Unlock()

// 	m.conns[conn.RemoteAddr()] = conn
// }

// func (m *Mux) Remove(addr net.Addr) {
// 	m.mu.Lock()

// 	defer m.mu.Unlock()

// 	delete(m.conns, addr)
// }

// func (m *Mux) SendMsg(msg string) error {
// 	m.mu.Lock()

// 	defer m.mu.Unlock()

// 	for _, conn := range m.conns {
// 		err := io.WriteString(conn, msg)

// 		if err != nil {
// 			return nil
// 		}
// 	}

// 	return nil
// }

// type Mux struct {
// 	add     chan net.Conn
// 	remove  chan net.Addr
// 	sendMsg chan string
// }

// func (m *Mux) Add(conn net.Conn) {
// 	m.add <- conn
// }

// func (m *Mux) Remove(addr net.Addr) {
// 	m.remove <- addr
// }

// func (m *Mux) SendMsg(msg string) error {
// 	m.sendMsg <- msg

// 	return nil
// }

// func (m *Mux) loop() {
// 	conns := make(map[net.Addr]net.Conn)

// 	for {
// 		select {
// 		case conn := <-m.add:
// 			conns[conn.RemoteAddr()] = conn

// 		case addr := <-m.remove:
// 			delete(conns, addr)

// 		case msg := <-m.sendMsg:
// 			for _, conn := range conns {
// 				io.WriteString(conn, msg)
// 			}
// 		}
// 	}
// }

type Mux struct {
	ops chan func(map[net.Addr]net.Conn)
}

func (m *Mux) Add(conn net.Conn) {
	m.ops <- func(m map[net.Addr]net.Conn) {
		m[conn.RemoteAddr()] = conn
	}
}

func (m *Mux) Remove(addr net.Addr) {
	m.ops <- func(m map[net.Addr]net.Conn) {
		delete(m, addr)
	}
}

func (m *Mux) SendMsg(msg string) error {
	m.ops <- func(m map[net.Addr]net.Conn) {
		for _, conn := range m {
			io.WriteString(conn, msg)
		}
	}

	return nil
}

func (m *Mux) loop() {
	conns := make(map[net.Addr]net.Conn)

	for _, op := range m.ops {
		op(conns)
	}
}
