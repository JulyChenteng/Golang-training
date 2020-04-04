package workers

import (
	"sync"
)

type Job struct {
	Data interface{}
	Proc func(interface{})
}

type Pool struct {
	works chan Job
	wg    sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		works: make(chan Job, 2048),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.works {
				w.Proc(w.Data)
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Job) {
	p.works <- w
}

func (p *Pool) Shutdown() {
	close(p.works)
	p.wg.Wait()
}
