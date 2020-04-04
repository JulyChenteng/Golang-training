package main

import (
	"sync"
	"fmt"
	"time"
)

var m sync.RWMutex

func f(n int) int {
	if n < 1 {
		return 0
	}
	fmt.Println("RLock")
	m.RLock()
	defer func() {
		fmt.Println("RUnlock")
		m.RUnlock()
	}()
	time.Sleep(100 * time.Millisecond)
	return f(n-1) + n
}

func main() {
	done := make(chan int)
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Lock")
		m.Lock()
		fmt.Println("Unlock")
		m.Unlock()
		done <- 1
	}()
	f(4)
	<-done
}