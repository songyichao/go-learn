package main

import "fmt"

func sayHello() {
	fmt.Println("hello , quanxiaoha.com ...")
}

func main() {
	// 声明一个函数类型的变量，注意类型为 func()
	var f func()
	// 将函数名赋值给变量 f
	f = sayHello
	// 通过变量 f 直接调用函数
	f()
}
