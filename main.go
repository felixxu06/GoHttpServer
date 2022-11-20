package main

import (
	"errors"
	"fmt"
)

func f1(num int) (int, error) {
	if 40 < num {
		return -1, errors.New("you are too old")
	}
	return 200, nil
}

type argError struct {
	arg  int
	prob string
}

func (arg *argError) Error() string {
	return fmt.Sprintf("%d - %s", arg.arg, arg.prob)
}

func f2(arg int) (int, error) {
	if arg > 42 {
		return -1, &argError{arg: arg, prob: "cann't work with it"}
	}

	return arg + 30, nil
}

func main() {
	numbers := []int{1, 2, 3, 4, 43, 55}
	for _, number := range numbers {
		if r, e := f1(number); e != nil {
			fmt.Println("f1 failed, ", e)
		} else {
			fmt.Println("f1 succeed", r)
		}
	}

	for _, number := range numbers {
		if r, e := f2(number); e != nil {
			fmt.Println("f2 failed, ", e)
		} else {
			fmt.Println("f2 succeed", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
