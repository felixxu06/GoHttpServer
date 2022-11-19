package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	permit() float64
}

type rect struct {
	height, weight float64
}

type cicle struct {
	redius float64
}

func (c cicle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c cicle) permit() float64 {
	return math.Pi * 2 * c.redius
}

func (r rect) area() float64 {
	return r.weight * r.height
}

func (r rect) permit() float64 {
	return r.height*2 + r.weight*2
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area of g is ", g.area())
	fmt.Println("permit of g is ", g.permit())
}

func main() {
	r := rect{height: 30, weight: 20}
	c := cicle{redius: 25}

	measure(r)
	measure(c)
}
