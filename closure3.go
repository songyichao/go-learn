package main

func main() {
	var dummy [3]int
	for i := 0; i < len(dummy); i++ {
		println(i)
	}

	var dummy1 [3]int
	var f func()
	for i := 0; i < len(dummy1); i++ {
		f = func() {
			println(i)
		}
	}
	f()

	var funcSlice []func()
	for i := 0; i < 3; i++ {
		funcSlice = append(funcSlice, func() {
			println(i)
		})

	}
	for j := 0; j < 3; j++ {
		funcSlice[j]()
	}
}
