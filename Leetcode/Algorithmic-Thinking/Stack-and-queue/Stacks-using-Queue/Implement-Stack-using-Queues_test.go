package Stacks_using_Queue

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Top())
}
