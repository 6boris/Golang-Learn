package main

import (
	"container/list"
	"fmt"
)

type MinStack struct {
	item list.List
	min  interface{}
}

/** initialize your data structure here. */
func Constructor() MinStack {
	res := MinStack{}
	return res
}

func (this *MinStack) Push(x interface{}) {
	this.item.PushBack(x)
	if this.min == nil || x.(int) < this.min.(int) {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	this.item.Remove(this.item.Front())
}

func (this *MinStack) Top() interface{} {
	return this.item.Front().Value
}

func (this *MinStack) GetMin() interface{} {
	return this.min
}

func (this *MinStack) lprint() {
	for e := this.item.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func main() {
	Min := Constructor()
	Min.Push(-2)
	Min.Push(-0)
	Min.Push(-3)
	fmt.Println(Min.GetMin())
	Min.lprint()

	Min.Pop()
	fmt.Println(Min.Top())
	Min.lprint()
	fmt.Println(Min.GetMin())

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
