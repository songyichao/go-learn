package main

import (
	"fmt"
	"unsafe"
)

type MyStruct1 struct {
	i int
	j int
}

func myFunction3(ms *MyStruct1) {
	ptr := unsafe.Pointer(ms)
	for i := 0; i < 2; i++ {
		c := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(8*i)))
		*c += i + 1
		fmt.Printf("[%p] %d\n", c, *c)
	}
}

func main() {
	a := &MyStruct1{i: 40, j: 50}
	myFunction3(a)
	fmt.Printf("[%p] %v\n", a, a)
}
