package main

import "fmt"

type Stack1 struct {
	data []int
}

func (self *Stack1) Pop() int {
	if len(self.data) <= 0 {
		fmt.Println("Stack is Empty")
		return -1
	}

	item := self.data[len(self.data)-1]
	self.data = self.data[:len(self.data)-1]

	return item
}

func (self *Stack1) Push(data int) {
	self.data = append(self.data, data)
}

func main() {
	s := new(Stack1)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.data)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
