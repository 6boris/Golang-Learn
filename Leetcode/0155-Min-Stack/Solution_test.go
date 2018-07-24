package Solution

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	stack := MinStack{}
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}
