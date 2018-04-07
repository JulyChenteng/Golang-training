package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", Cookie2)
	http.ListenAndServe(":8081", nil)
}

func Cookie(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}

	http.SetCookie(w, ck)
	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, ck2.Value)
}

func Cookie2(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello world",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}

	w.Header().Set("Set-Cookie", strings.Replace(ck.String(), " ", "%20", -1))
	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, ck2.Value)
}
