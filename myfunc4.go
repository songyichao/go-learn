package main

type MyStruct2 struct {
	i int
	j int
}

func myFunction4(ms *MyStruct2) *MyStruct2 {
	return ms
}

func main() {
	myFunction4(&MyStruct2{i: 40, j: 50})
}
