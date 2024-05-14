package main

import "fmt"

func ret(a int) int {
	return a
}

func ret1(a int) (b int) {
	b = a
	return
}

func main() {
	fmt.Println(ret(2))
	fmt.Println(ret1(2))
}
