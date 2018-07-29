package Solution

import "container/list"

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
func (s *Stack) Push(x interface{}) {
	s.data.PushBack(x)
}

//	Return the stack top val
func (s *Stack) Top() interface{} {
	return s.data.Back().Value
}

//	Delete and Return the stack top val
func (s *Stack) Pop() interface{} {
	res := s.data.Back().Value
	s.data.Remove(s.data.Back())
	return res
}
