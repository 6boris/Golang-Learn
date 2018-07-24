package Solution

import (
	"errors"
)

//	1.将栈定义为一个借口方便插入任意类型数据
//	2.采用切片可以动态分配大小。还有Len() 和Cap()这个2个战术可以免费使用
type Stack []interface{}

//	返回栈现在元素个数
func (stack Stack) Len() int {
	return len(stack)
}

//	返回栈的容量
func (stack Stack) Cap() int {
	return cap(stack)
}

//	判断栈是否为空
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
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

type MinStack struct {
	data Stack
	min  Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x interface{}) {
	this.data.Push(x)

	if n, _ := this.min.Top(); this.min.IsEmpty() || n.(int) >= x.(int) {
		this.min.Push(x)
	}
}

func (this *MinStack) Pop() {
	x, _ := this.data.Pop()

	if n, _ := this.min.Top(); n == x {
		this.min.Pop()
	}
}

func (this *MinStack) Top() interface{} {
	res, _ := this.data.Top()
	return res
}

func (this *MinStack) GetMin() interface{} {
	res, _ := this.min.Top()
	return res
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
