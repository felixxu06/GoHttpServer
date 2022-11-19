package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	name string
}

func main() {
	co := container{
		base: base{
			num: 3,
		},
		name: "cicle",
	}

	fmt.Printf("container ={num: %v, name: %v}", co.base.num, co.name)

	fmt.Println("also num: %v", co.base.num)

	fmt.Println("desrcibe ", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co

	desc := d.describe()

	fmt.Println(desc)
}
