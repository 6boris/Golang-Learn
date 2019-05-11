/*
go run -gcflags '-m -l' escape2.go

x没有逃逸

identity 函数直接当成返回值了因为没有对z做引用，
所以z没有逃逸。
对x的引用也没有逃出main函数的作用域，因此x也没有发送逃逸

*/

package main

type S2 struct{}

func main() {
	var x S2
	y := &x
	_ = *identity2(y)
}
func identity2(z *S2) *S2 {
	return z
}
