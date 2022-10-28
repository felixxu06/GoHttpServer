package main

import "net/http"
import "fmt"
import "io"


func main()  {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/greeting", greetingHanlder)
	http.ListenAndServe(":3333", nil)
}

func rootHandler(resp http.ResponseWriter, req *http.Request)  {
	fmt.Printf("root handler invoked")
	io.WriteString(resp, "Welcome to my website")
}

func greetingHanlder(resp http.ResponseWriter, req *http.Request)  {
	fmt.Printf("greeting handler invoked")
	io.WriteString(resp, "hello visitor, hoping you have a good day")
}