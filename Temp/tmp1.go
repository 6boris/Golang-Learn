package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	//	定时直接关闭channel
	go func() {
		time.Sleep(time.Second * 3)
		close(jobs)
	}()

	//	无脑向channel发送数据【向已经关闭的channel发送数据会直接panic】
	go func() {
		for i := 0; ; i++ {
			jobs <- i
			fmt.Println("produce:", i)
		}
	}()
	//	在channel中读取数据
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range jobs {
			fmt.Println("consume:", i)
		}
	}()
	wg.Wait()
}
