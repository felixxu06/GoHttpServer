package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}
