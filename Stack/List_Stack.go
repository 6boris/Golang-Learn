package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	data list.List
}

func (self *Stack) PrintList() {
	for e := self.data.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func (self *Stack) Pop() int {
	if self.data.Len() == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	//吃有问题待解决
	//item := self.data.Back().Value.(int)

	item := fmt.Sprintf("%d", self.data.Front().Value)
	self.data.Remove(self.data.Back())

	return item
}

func (self *Stack) Push(data int) {
	self.data.PushBack(data)
}

func main() {
	stk := Stack{}
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Push(4)
	stk.PrintList()
	stk.Pop()
	stk.PrintList()
}
