package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = []byte("123")
	s := *(*string)(unsafe.Pointer(&b)) //A

	b[1] = '4'
	fmt.Println(b)
	fmt.Println(s)
}
