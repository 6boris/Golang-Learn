/*
go run -gcflags '-m -l' escape4.go

y 逃逸
refStruct4 函数对Y取了引用，所以y发送的逃逸

*/
package main

type S4 struct {
	M *int
}

func main() {
	var i int
	refStruct4(i)
}

func refStruct4(y int) (z S4) {
	z.M = &y
	return z
}
