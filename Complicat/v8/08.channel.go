package main

import "fmt"

func main() {
	var ch chan int
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}

/**
A: 不能编译
B: 输出 1
C: 输出 0
D: panic
*/
