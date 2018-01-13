package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayhello)

	wd, err := os.Getwd() //获取工作目录
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir(wd)))) //实现简易静态文件服务器

	log.Fatal(http.ListenAndServe(":8080", mux))
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL:"+r.URL.String())
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is version 2.0")
}
