// func Query(conns []Conn, query string) Result {
// 	ch := make(chan Result, len(conns))  // buffered

// 	for _, conn := range conns {
// 		go func(c Conn) {
// 			ch <- c.DoQuery(qurey):
// 		}(conn)
// 	}

// 	return <-ch
// }
