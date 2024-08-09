package main

import (
	"errors"
	"fmt"
)

func main() {
	e1()
	e2()
	e0()
}

func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("defer1 error")
	return
}
func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("defer2 error")
	return
}
func e0() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("defer0 error")
	return
}
