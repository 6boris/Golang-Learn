package main

import "fmt"

type MyStruct struct {
	i int
}

func myFunction(a MyStruct, b *MyStruct) {
	a.i = 31
	b.i = 41
	fmt.Printf("in my_function - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}

func main() {
	a := MyStruct{i: 30}
	b := &MyStruct{i: 40}
	fmt.Printf("before calling - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
	myFunction(a, b)
	fmt.Printf("after calling  - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}

//$ go run main.go
//before calling - a=({30}, 0xc000018178) b=(&{40}, 0xc00000c028)
//in my_function - a=({31}, 0xc000018198) b=(&{41}, 0xc00000c038)
//after calling  - a=({30}, 0xc000018178) b=(&{41}, 0xc00000c028)
