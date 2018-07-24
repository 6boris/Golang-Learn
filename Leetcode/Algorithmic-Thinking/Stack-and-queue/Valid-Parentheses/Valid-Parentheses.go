package main

import (
	"container/list"
	"fmt"
)

//	自己写的一个栈，注意interface传递参数的转换
type Stack struct {
	data list.List
}

func (this *Stack) Len() int {
	return this.data.Len()
}

func (this *Stack) IsEmpty() bool {
	return this.data.Len() == 0
}

func (this *Stack) Push(x interface{}) {
	this.data.PushBack(x)
}

func (this *Stack) Pop() interface{} {
	if this.data.Len() == 0 {
		return nil
	} else {
		res := this.data.Back().Value
		this.data.Remove(this.data.Back())
		return res
	}
}

func (this *Stack) Top() interface{} {
	if this.data.Len() == 0 {
		return nil
	} else {
		return this.data.Back().Value
	}
}

func isValid(s string) bool {
	stack := Stack{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack.Push(s[i])
		} else {
			if stack.IsEmpty() {
				return false
			}
			//	这是是不能直接将stack.top 和 ‘(’相比较的 ,因为string 类型的单个字符经过interface转换之后会变为uint8
			if s[i] == ')' && stack.Top() != uint8('(') {
				return false
			}
			if s[i] == ']' && stack.Top() != uint8('[') {
				return false
			}
			if s[i] == '}' && stack.Top() != uint8('{') {
				return false
			}
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

func main() {

	//	40 (
	//	41 )
	//	91	[
	//	93	]
	//	123	{
	//	125 }

	str := "(]"
	fmt.Println(isValid(str))
}
