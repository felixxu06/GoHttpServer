package main

import (
		"net/http"
		"fmt"
		"io"
		"elmo.com/core"
	)

func main()  {
	var server core.GoRouter = new(core.Routers)
	server.Register("/", rootHandler)
	server.Register("/greeting", greetingHanlder)

	server.Start(3333)
}

func rootHandler(resp http.ResponseWriter, req *http.Request)  {
	fmt.Printf("root handler invoked")
	io.WriteString(resp, "Welcome to my website")
}

func greetingHanlder(resp http.ResponseWriter, req *http.Request)  {
	fmt.Printf("greeting handler invoked")
	io.WriteString(resp, "hello visitor, hoping you have a good day")
}