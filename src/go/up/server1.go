package main

import (
	. "fmt"
	"log"
	"net/http"
)

func main() {
	// 搞一个简单的web服务器
	// http://localhost:8000/hello?q=1001
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// http头部信息
	Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	Fprintf(w, "Host = %q\n", r.Host)
	Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
