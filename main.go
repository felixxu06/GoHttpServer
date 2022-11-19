package main

import "fmt"

type person struct {
	age  int
	name string
}

func newPerson(name string) *person {
	p := person{
		name: name,
	}
	p.age = 37

	return &p
}

func main() {
	fmt.Println(person{35, "Anita"})
	fmt.Println(person{name: "Xinxin", age: 7})
	fmt.Println(person{name: "Xincheng"})

	fmt.Println(&person{name: "New Zealand", age: 100})
	fmt.Println(newPerson("Nobody"))

	s := person{name: "China", age: 5000}
	fmt.Println(s.name)

	sp := s

	sp.age = 300
	fmt.Println(s.age)
}
