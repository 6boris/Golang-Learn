package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		nums := []int{1, 2, 2, 3, 3, 4, 5}
		data := MakeList(nums)
		fmt.Println("Source:")
		PrintList(data)
		got := deleteDuplicates(data)
		fmt.Println("Reverse:")
		PrintList(got)
	})
}
