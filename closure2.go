package main

func main() {
	x := 1
	f := func() {
		println(x)
	}
	x = 2
	x = 3
	f()

	y := 1
	func() {
		println(y)
	}()
	y = 2
	y = 3
}
