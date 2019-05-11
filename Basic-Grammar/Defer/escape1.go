/*
go run -gcflags '-m -l' escape1.go

没有逃逸

值传递，直接在栈上分配

Go语言函数传递都是通过值传递的，调用的时候直接在栈上copy出一份参数，不存在逃逸

*/

package main

type S1 struct{}

func main() {
	var x S1
	_ = identity1(x)
}
func identity1(x S1) S1 {
	return x
}
