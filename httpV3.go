package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {

	//自定义Server
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	//创建路由管理的结构，以及注册路由和相应的处理函数
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/bye"] = saybye
	mux["/hello"] = sayhello

	log.Fatal(server.ListenAndServe())
}

//自定义Server的请求处理函数
type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "URL:"+r.URL.String())
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello,this is version 3.0")
}

func saybye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye, this is version 3.0")
}
