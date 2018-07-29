package Solution

import (
	"container/list"
	"errors"
)

// define a new type Stack
type Stack struct {
	data list.List
}

//	Return the stack len
func (s *Stack) Len() int {
	return s.data.Len()
}

//	Judge the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.data.Len() == 0
}

//	Push a val to stack top
func (s *Stack) Push(x int) {
	s.data.PushBack(x)
}

//	Return the stack top val
func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Out of index, len is 0")
	}
	return s.data.Back().Value, nil
}

//	Delete and Return the stack top val
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Out of index, len is 0")
	}
	res := s.data.Back().Value.(int)
	s.data.Remove(s.data.Back())
	return res, nil
}

func dailyTemperatures(temperatures []int) []int {
	dist := []int{}
	index := Stack{}

	for curIndex := 0; curIndex < len(temperatures); curIndex++ {

		index.Push(curIndex)
	}
	return dist
}
