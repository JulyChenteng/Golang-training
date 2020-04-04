package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-ch)
	}
}

/*
func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}
*/

/*
func createWorker(id int) chan int {
	ch := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-ch)
		}
	} ()

	return ch
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}
*/

/*
func createWorker(id int) chan int {
	ch := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-ch)
		}
	} ()

	return ch
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}
*/

/*
func createWorker(id int) chan<- int {
	ch := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-ch)
		}
	} ()

	return ch
}

func chanDemo() {
	var channels [10]chan<- int //该通道只能用于发送
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}
*/

func bufferedChannel() {
	c := make(chan int, 3) // 缓存大小为3
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func worker2(id int, c chan int) {
	/*
		//第一种方法
		for {
			ch, ok := <-c
			if !ok {   //数据是否发送完
				break
			}

			fmt.Printf("worker %d received %c\n", id, ch)
		}
	*/
	//第二种方法
	for ch := range c {
		fmt.Printf("worker %d received %c\n", id, ch)
	}
}

func channelClose() {
	c := make(chan int, 3) // 缓存大小为3
	go worker2(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c) //发送完后关闭通道
	time.Sleep(time.Millisecond)
}

func main() {
	channelClose()
	//bufferedChannel()
	//chanDemo()
}
