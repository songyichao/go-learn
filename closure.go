package main

import "fmt"

// 定义一个累加函数，返回类型为 func() int, 入参为整数类型，每次调用函数对该值进行累加
func Add(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回累加值
		return value
	}
}

func main() {
	// 创建一个累加器，初始值为 1
	accumulator := Add(1)

	// 累加1并打印
	fmt.Println(accumulator())
	// 再来一次
	fmt.Println(accumulator())

	// 创建另一个累加器，初始值为 10
	accumulator2 := Add(10)

	// 累加1并打印
	fmt.Println(accumulator2())
}
