package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)

	p := fmt.Println

	if err != nil {
		panic(err)
	}

	p(u.Scheme)
	p(u.User)
	p(u.User.Username())

	pwd, _ := u.User.Password()
	p(pwd)
	p(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)

	p(host)
	p(port)

	p(u.Path)
	p(u.Fragment)

	p(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])

}
