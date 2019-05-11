package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	//wg.Done()
	wg.Wait()
}

/**
A: 不能编译
B: 无输出，正常退出
C: 程序hang住
D: panic
*/
