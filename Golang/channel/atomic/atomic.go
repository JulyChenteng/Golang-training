package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	mutex sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() {
		a.mutex.Lock()
		defer a.mutex.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increment()

	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
