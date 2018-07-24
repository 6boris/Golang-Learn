package Queue_using_Stacks

//package main

import (
	"errors"
)

//	232
//	地址：https://leetcode.com/problems/implement-queue-using-stacks/description/
//	注意：因为Golang自身是没有栈这个数据结构的所以，我们需要自己写一个栈

//	1.将栈定义为一个借口方便插入任意类型数据
//	2.采用切片可以动态分配大小。还有Len() 和Cap()这个2个战术可以免费使用
type Stack []interface{}

//	返回栈现在元素个数
func (stack Stack) Len() int {
	return len(stack)
}

//	判断栈是否为空
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

//	返回栈的容量
func (stack Stack) Cap() int {
	return cap(stack)
}

//	元素入栈
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

//	返回栈顶元素，如果栈为空则元素为空error不为空
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

//	元素出栈，如果栈为空则元素为空error不为空
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

type MyQueue struct {
	//	in 用来存放入栈元素
	in Stack
	//	out 用来存放出栈元素
	out Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x interface{}) {
	this.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() interface{} {
	if this.out.IsEmpty() {
		for !this.in.IsEmpty() {
			if res, _ := this.in.Pop(); res != nil {
				this.out.Push(res)
			}
		}
	}
	res, _ := this.out.Pop()
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() interface{} {
	if this.out.IsEmpty() {
		for !this.in.IsEmpty() {
			if res, _ := this.in.Pop(); res != nil {
				this.out.Push(res)
			}
		}
	}
	res, _ := this.out.Top()
	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.in.IsEmpty() && this.out.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
