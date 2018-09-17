package workerpool

const (
	MAXWORKERS = 256
)

type Job struct {
	Data interface{}
	Proc func(interface{})
}

var (
	once sync.Once
	pool *WorkerPool
)

type WorkerPool struct {
	works chan Job
}

func WorkerPoolInstance() *WorkerPool {
	once.Do(func() {
		pool = newWorkerPool(MAXWORKERS)
	})

	return pool
}

func newWorkerPool(maxGoroutines int) *WorkerPool {
	pool := WorkerPool{
		works: make(chan Job),
	}

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range pool.works {
				w.Proc(w.Conn, w.Data)
			}
		}()
	}

	return &pool
}

func (p *WorkerPool) Run(w Job) {
	p.works <- w
}

func (p *WorkerPool) Shutdown() {
	close(p.works)
}
