package main

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}
func main() {
	go f()
	c <- 0
	print(a)
}

/**
A: 不能编译
B: 输出 1
C: 输出 0
D: panic
*/
