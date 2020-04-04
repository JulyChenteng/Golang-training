package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println("hello from goroutine ", i)
			}
		}(i)
	}

	time.Sleep(5 * time.Millisecond)
}
