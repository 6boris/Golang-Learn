package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		stack := Stack{}
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)

		fmt.Println(stack.Pop())
		fmt.Println(stack.Top())
		fmt.Println(stack.Top())

	})
	t.Run("Test-2", func(t *testing.T) {
		fmt.Println("demo")
	})

}
