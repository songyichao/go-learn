package main

func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	i := incr()
	println(i())
	println(i())
	println(i())

	println(incr()())
	println(incr()())
	println(incr()())
}
