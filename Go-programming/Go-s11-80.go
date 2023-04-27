func First(query string, replicas ...Search) Result {
	c := make(chan Result)

	searchReplica := func(i int) {
		c <- replicas[i](query)
	}

	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	result := First("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))

	elapsed := time.Since(start)

	fmt.Println(result)
	fmt.Println(elapsed)
}

c := make(chan Result)

go func() {
	c <- First(query, Web1, Web2)
}()

go func() {
	c <- First(query, Image1, Image2)
}()

go func() {
	c <- First(query, Video1, Video2)
}

timeout := time.After(80 * time.Millisecond)

for i := 0; i < 3; i++ {
	select {
	case result := <-c:
		results = append(results, result)

	case <-timeout:
		fmt.Println("timed out")

		return
	}
}

return
