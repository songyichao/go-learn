package main

import "fmt"

func main() {
	result := modifyNamedReturnValue()
	fmt.Println(result)
}

func modifyNamedReturnValue() (result int) {
	result = 1
	defer func() {
		result = 2
	}()
	return 0
}
