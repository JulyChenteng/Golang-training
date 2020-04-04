package main

import (
	"fmt"
	"time"
	//"math/rand"
	"math/rand"
)

func worker(id int, ch chan int) {
	for n := range ch {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)

	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)

			out <- i
			i++
		}
	}()

	return out
}

func main() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)
	/*
		for {
			select {
			case n := <-c1:
				fmt.Println("received from c1", n)
			case n := <-c2:
				fmt.Println("received from c2:", n)
				//default:
				//	fmt.Println("no value received!")
			}
		}
	*/

	/*
		for {
			select {
			case n := <-c1:
				w <- n               //弊端：w会阻塞到数据被读取才会继续执行
			case n := <-c2:
				w <- n
				//default:
				//	fmt.Println("no value received!")
			}
		}
	*/

	/*
		n := 0
		hasValue := false
		for {
			var activeWorker chan<- int
			if hasValue {
				activeWorker = w
			}

			select {
				case n = <-c1:                //弊端：如果在worker端读取之后处理数据速度太慢，新的n的值在未进入activeWorker之前会被冲掉
					hasValue = true
				case n = <-c2:
					hasValue = true
				case activeWorker <- n:      //数据没有准备好之前activeWorker为nil chan，在写入数据时会阻塞，如果不是nil chan，
		那么会在数据未准备好之前，读入脏数据。
					hasValue = false
				//default:
				//	fmt.Println("no value received!")
			}
		}
	*/
	/*
		n := 0
		var values []int
		for {
			var activeWorker chan<- int
			var activeValue int
			if len(values) > 0 {
				activeWorker = w
				activeValue = values[0]
			}

			select {
			case n = <-c1:
				values = append(values, n)         //将传入的数据先缓存起来
			case n = <-c2:
				values = append(values, n)
			case activeWorker <- activeValue:
				values = values[1:]
				//default:
				//	fmt.Println("no value received!")
			}
		}
	*/

	//十秒钟停止程序
	n := 0
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second) //定时器
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}

		select {
		case n = <-c1:
			values = append(values, n) //将传入的数据先缓存起来
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]

		case <-tick:
			fmt.Println("queue len = ", len(values)) //定时器实现每一秒查看缓存队列长度
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tm:
			fmt.Println("bye") //如果tm通道有数据，则退出
			return
		}
	}
}
