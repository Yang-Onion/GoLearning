package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {

	if arg == 0 {
		return -1, errors.New("can't work with 0")
	}

	return arg, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s ", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 0 {
		return -1, &argError{arg: -1, prob: "can't work with 0"}
	}

	return arg, nil
}

func funcError() {
	for _, i := range []int{-1, 0, 1} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{-1, 0, 1} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(0)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
