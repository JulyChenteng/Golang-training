package service

import (
	"sync"

	"../../utility/message"
	"../defs"
)

const (
	MAXWORKERS = 256
)

type Job struct {
	Data message.Message
	Conn defs.Conn
	Proc func(conn defs.Conn, message message.Message)
}

func init() {
	Pool = WorkerPoolInstance()
}

var (
	once sync.Once
	Pool *WorkerPool
)

type WorkerPool struct {
	works chan Job
}

func WorkerPoolInstance() *WorkerPool {
	var pool *WorkerPool
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
	once.Do(func() {
		close(p.works)
	})
}
