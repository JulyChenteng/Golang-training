package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//设置路由
	http.HandleFunc("/", sayhello)

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal("error")
	}
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is version 1.0")
}
