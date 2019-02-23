package main

import (
	"fmt"
	"net"
	"sync"
)

var gameWait *sync.WaitGroup//此处也申明为指针变量
//参数传递为指针传递*sync.WaitGroup，即wait 指向的是main里的waitGroup的内存地址
func ConnSocket(serverAddr string, wait *sync.WaitGroup) {
	var err error
	Conn, err = net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connected:", err.Error())
		wait.Done()
		return
	}

	gameWait = wait//指针指向另一个指针，这样就做到了gameWait指向的是wait的地址，所以gameWait修改wait的内存地址锁存贮的值也会跟随改变

	fmt.Println("connected OK:", Conn.RemoteAddr())
	go readMessage()
}

func readMessage() {
	fmt.Println("readMessage:")
	for {
		if OK:=dosomething(),OK{
			gameWait.Done()
		}
	}
}
