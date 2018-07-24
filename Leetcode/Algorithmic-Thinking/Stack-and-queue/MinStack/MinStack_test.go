package main

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	MinStack := Constructor()
	MinStack.Push(1)
	MinStack.Push(2)

	fmt.Println(MinStack)

	fmt.Println(MinStack.Top())
	fmt.Println(MinStack.GetMin())
	fmt.Println(MinStack)

	MinStack.Pop()
	fmt.Println(MinStack.GetMin())

	fmt.Println(MinStack.Top())

	fmt.Println(MinStack)

}
