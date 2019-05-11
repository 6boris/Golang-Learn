/*
go run -gcflags '-m -l' escape5.go

i 没有发送逃逸

在main函数离对i取了引用并且把它传给refStruct函数，i的引用一直在main函数的作用域使用，因此i没有发生逃逸。

和上一个例子相比，i的写法有一点小差别，但是i的命运是不同的，导致程序小效果是不同的：
	例4中i在main函数的作用域栈中分配
	本例中i只分配了一次，然后通过引用传递。

*/
package main

type S5 struct {
	M *int
}

func main() {
	var i int
	refStruct5(&i)
}

func refStruct5(y *int) (z S5) {
	z.M = y
	return z
}
