package main

import (
	"fmt"

	"./mock"
	"./real"
	"time"
	"unsafe"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = &mock.Retriever{"This is a mock retriever!"}
	fmt.Printf("%T, %v\n", r, r)

	fmt.Println(download(r))
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Println(download(r))
	fmt.Printf("%T, %v\n", r, r)
	inspect(r)

	//Type assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.UserAgent)
	} else {
		fmt.Println("not a real retriever!")
	}

	fmt.Println(unsafe.Sizeof(r))                            //16
	fmt.Println(unsafe.Sizeof(mock.Retriever{"helloworld"})) //16
	fmt.Println(unsafe.Sizeof(real.Retriever{}))             //24

	fmt.Println(unsafe.Sizeof(string("")))               //16
	fmt.Println(unsafe.Sizeof(real.Retriever{}.TimeOut)) //8
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents", v.Contents)
	case *mock.Retriever:
		fmt.Println(" Pointer Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}
