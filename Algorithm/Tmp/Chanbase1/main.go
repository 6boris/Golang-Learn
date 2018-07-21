package main

import (
	"fmt"
	"time"
)

/*
	Sent: a [Sender]
	Sent: b [Sender]
	Sent: c [Sender]
	Sent a sync singal... [Sender]
	Receive a sync singal and wait a second ...[Receiver]
	Received: a [Receiver]
	Received: b [Receiver]
	Received: c [Receiver]
	Sent: d [Sender]
	Received: d [Receiver]
	Wait 2 Seconds ... [Sender]
	Stopped...[Receiver]
*/

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	// 用于演示接收操作
	go func() {
		<-syncChan1
		fmt.Println("Receive a sync singal and wait a second ...[Receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[Receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped...[Receiver]")
		syncChan2 <- struct{}{}
	}()

	// 用于演示发送操作
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[Sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync singal... [Sender]")
			}
		}
		fmt.Println("Wait 2 Seconds ... [Sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2

}
