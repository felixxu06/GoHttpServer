package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
	val        T
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) GetAll() []T {
	var elem []T
	for e := list.head; e != nil; e = e.next {
		elem = append(elem, e.val)
	}
	return elem
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys", MapKeys(m))

	_ = MapKeys[int, string](m)
	list := List[int]{}
	list.Push(10)
	list.Push(20)
	list.Push(35)
	fmt.Println("list", list.GetAll())
}
