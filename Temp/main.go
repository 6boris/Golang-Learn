package main

import (
	"fmt"
	"github.com/kylesliu/gonp/Stack"
)

func main() {
	s := Stack.GetMinStack()
	fmt.Println(s.Top())
	fmt.Println(s.IsEmpty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Top())
}
