package main

import "fmt"

func main() {
	x := new(int)
	*x = 10
	fmt.Println(*x)
}

func printInt(i *int) {
	fmt.Println(*i)
}
