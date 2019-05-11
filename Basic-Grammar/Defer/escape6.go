/**
go run -gcflags '-m -l' escape6.go

i 逃逸，x未逃逸，x的作用域一直在main函数中


本列i发生了逃逸，按照前面例子5的分析，i不会逃逸，但是此列中i赋值给了一个输入的struct，go不能从输出反推到输入
两个例子的区别是例子5中的S是在返回值字的，输入只能流入到输出。
本列中的S是在输入参数中，所以逃逸分析失败，i要逃逸到堆山。

*/
package main

type S6 struct {
	M *int
}

func main() {
	var x S6
	var i int
	refStruct6(&i, &x)
}

func refStruct6(y *int, z *S6) {
	z.M = y
}
