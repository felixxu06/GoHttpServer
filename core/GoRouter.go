package core

import ("net/http")

type GoRouter interface {
	
	Register(path string, handler func(http.ResponseWriter, *http.Request))

	FindRoute(path string) (handler func(http.ResponseWriter, *http.Request), exist bool)

	Start(port int) error
}