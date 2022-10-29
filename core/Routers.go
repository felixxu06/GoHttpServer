package core

import (
	"fmt"
	"net/http"
)

type Routers struct {
	internalRouters map[string]func(http.ResponseWriter, *http.Request)
}

func (route *Routers) Register(path string, handler func(http.ResponseWriter, *http.Request)) {
	route.internalRouters[path] = handler
}

func (route *Routers) FindRoute(path string) func(http.ResponseWriter, *http.Request) {
	return route.internalRouters[path]
}

func (route *Routers) Start(port int) error  {
	
	for key, value := range route.internalRouters {
        http.HandleFunc(key, value)
    }

   return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}