package main

import (
"fmt"
"pressure/game"
"sync"
)

var waitGroup sync.WaitGroup //定义一个同步等待的组
func main(){
waitGroup.Add(1) //添加一个计数
//这里传递waitGroup的内存地址
go game.ConnSocket(serverAddr, &waitGroup) //调用其他包的方法执行任务
waitGroup.Wait() //阻塞直到所有任务完成

fmt.Println("main DONE!!!")
}
