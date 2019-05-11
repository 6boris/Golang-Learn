/*
go run -gcflags '-m -l' escape3.go

z 发生逃逸
z 是对x的拷贝，ref函数中对z取了引用,所以不能放在栈上，通过引用不能找到z，所以z必须要逃逸到堆山。
尽管在main函数中，直接丢弃了ref的结果，但是go的编译器寒梅与那么智能，分析不出来这种情况。
而对x从来就没有去引用，所以x不会发生逃逸.
*/
package main

type S3 struct{}

func main() {
	var x S3
	_ = *ref3(x)
}
func ref3(z S3) *S3 {
	return &z
}
