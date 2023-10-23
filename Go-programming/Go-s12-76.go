go func() { // copy input to output
	for val := range input {
		output <- val
	}
}()

type Work struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.z * w.y
		Sleep(w.z)

		out <- w
	}
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)

	for i := 0; i < NumWorkers; i++ {
		go worker(in, out)
	}

	go sendLotsOfWork(in)
	receiveLotsOfResults(out)
}

type Request struct {
	fn func() int // The operation to perform.
	c  chan   int // The channel to return the result.
}

func requster(work chan<- Request) {
	c := make(chan int)

	for {
		// Kill some time (fake load).
		Sleep(rand.Int63n(nWoker * 2 * Second))
		work <- Request{workFn, c} // send request
		result := <-c              // wait for answer
		furtherProcess(resultb)
	}
}

type Worker struct {
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.request // get Request from balancer
		req.c <- req.fn()  // call fn and send result
		done  <- w         // we've finished this request
	}
}

type Pool []*Worker

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work:     // received a Request...
			b.dispatch(req) // ...so send it to a Worker

		case w := <-b.done:     // a worker has finished...
			b.completed(w)  // ...so update its info
		}
	}
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

// Send Request to worker
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task
	w.requests <- req
	// One more in its work queue.
	w.pending++
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.pending--
	// Remove it from heap.
	heap.Remove(&b.pool, w.index)
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}
