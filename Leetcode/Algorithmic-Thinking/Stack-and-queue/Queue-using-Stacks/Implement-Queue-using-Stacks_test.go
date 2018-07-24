package Queue_using_Stacks

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)

	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
