// 本题为考试单行多行输入输出规范示例，无需提交，不计分。
package main

import (
	"container/list"
	"fmt"
)

//   . 46
//	/ 47
func main() {

	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//str := input.Text()

	str := "/home/"
	fmt.Println(str)
	stk1 := NewStack()

	for i := 0; i < len(str); i++ {
		stk1.Push(str[i])
	}
	for i := 0; i < len(str); i++ {
		tmp := stk1.Pop()
		if isPie(tmp.(byte)) && isStr(stk1.Pop().(byte)) {
			continue
		}

	}

}

func isStr(b byte) bool {
	if b >= 97 && b <= 122 || b >= 65 && b <= 90 {
		return true
	}
	return false
}
func isDian(b byte) bool {
	if b == byte('.') {
		return true
	}
	return false
}

func isPie(b byte) bool {
	if b == byte('/') {
		return true
	}
	return false
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
