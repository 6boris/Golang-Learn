package main

import (
	"container/list"
	"fmt"
)

//	因为Golang 自带一个list和Queue已经非常类似，我只是简单的封装了一下

type Queue struct {
	item list.List
}

func (queue *Queue) Len() int {
	return queue.item.Len()
}

//	因为是链表所以这个方法没有意义
func (queue *Queue) Cap() int {
	return queue.item.Len()
}

func (queue *Queue) IsEmpty() bool {
	return queue.item.Len() == 0
}

//	向队列末尾添加元素
func (queue *Queue) Push(x interface{}) {
	queue.item.PushBack(x)
}

//	弹出队列首部的元素
func (queue *Queue) Pop() interface{} {
	res := queue.item.Front().Value
	queue.item.Remove(queue.item.Front())
	return res
}

//	获取队列首部元素，不弹出
func (queue *Queue) Peek() interface{} {
	return queue.item.Front().Value
}

func main() {
	queue := Queue{}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	fmt.Println(queue.Pop())

	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())

	//l := list.New()

}

//	这是一个失败的案列	在POP的时候会有问题
//type Queue interface {}
//
//func (queue *Queue) Len() int {
//	return len(*queue)
//}
//
//func (queue *Queue) Cap() int {
//	return cap(*queue)
//}
//
//func (queue *Queue) IsEmpty() bool {
//	return len(*queue) == 0
//}
//
////	向队列末尾添加元素
//func (queue *Queue) Push(x interface{}) {
//	*queue = append(*queue, x)
//}
//
////	弹出队列首部的元素
//func (queue *Queue) Pop() interface{} {
//	res := (*queue)[0]
//	res re
//	*queue = append((*queue)[:0], (*queue)[1:])
//	fmt.Println(queue)
//	return res
//}
//
////	获取队列首部元素，不弹出
//func (queue *Queue) Peek() interface{} {
//	res := (*queue)[0]
//	*queue = append((*queue)[:0], (*queue)[1:])
//	return res
//}
