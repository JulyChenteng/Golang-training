package main

import (
	"fmt"
	"time"
	"unsafe"
)

type i interface{}
type b struct{
	Name string
}

func (b2 b) Get() string {
	return b2.Name
}

func main() {
	ch := make(chan string)

	go sendData(ch)
	getData(ch)

	time.Sleep(1e9)


	var c i
	fmt.Println(unsafe.Sizeof(i))
	c = &b{Name : "hello"}
	fmt.Println(c.Name)
	fmt.Println(c.Name)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	//for {
		input = <-ch
		fmt.Printf("%s ", input)
	//}
}


