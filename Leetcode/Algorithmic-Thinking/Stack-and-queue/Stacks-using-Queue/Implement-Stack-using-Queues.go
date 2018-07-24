package Stacks_using_Queue

import (
	"container/list"
)

// 用栈实现队列 https://leetcode.com/problems/implement-stack-using-queues/description/
type MyStack struct {
	item list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	res := MyStack{}
	return res
}

/** Push element x onto stack. */
func (this *MyStack) Push(x interface{}) {
	this.item.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() interface{} {
	if this.item.Len() == 0 {
		return nil
	}
	res := this.item.Back().Value
	this.item.Remove(this.item.Back())
	return res
}

/** Get the top element. */
func (this *MyStack) Top() interface{} {
	if this.item.Len() == 0 {
		return nil
	}
	return this.item.Back().Value
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.item.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
