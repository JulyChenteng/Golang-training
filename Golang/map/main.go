package main

import (
	"./safemap"
	"fmt"
	"sync"
)

func main() {
	testOriginalMap()

	//testSafeMap()

	//testSyncMap()
}

func testOriginalMap() {
	mp := make(map[int]int)
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mp[j] = j
			}
		}()
	}
	wg.Wait()

	for k, v := range mp {
		fmt.Println(k, v)
	}
	fmt.Println()
}

func testSafeMap() {
	smp := safemap.Initial()
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				smp.Write(j, j)
			}
		}()
	}
	wg.Wait()

	smp.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	fmt.Println()
}

func testSyncMap() {
	smp := sync.Map{}
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				smp.Store(j, j)
			}
		}()
	}
	wg.Wait()

	smp.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
