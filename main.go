package main

import (
		"net/http"
		"fmt"
		"io"
	)

type GoRouter interface {

	Register(path string, handler func(http.ResponseWriter, *http.Request))

	// FindRoute(path string) (handler func(http.ResponseWriter, *http.Request), exist bool)

	Start(port int) error
}



type Routers struct {
	internalRouters map[string]func(http.ResponseWriter, *http.Request)
}

func (route *Routers) Register(path string, handler func(http.ResponseWriter, *http.Request)) {
	route.internalRouters[path] = handler
}

// func (route *Routers) FindRoute(path string) func(http.ResponseWriter, *http.Request) {
// 	return route.internalRouters[path]
// }

func (route *Routers) Start(port int) error  {

	for key,     value := range route.internalRouters {
        http.HandleFunc(key, value)
	}

   return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main()  {
	var server GoRouter = &Routers{internalRouters: make(map[string]func(http.ResponseWriter, *http.Request), 10)}
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