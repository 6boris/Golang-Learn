package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		data := MakeList(nums)
		fmt.Println("Source:")
		PrintList(data)
		got := swapPairs(data)
		fmt.Println("Reverse:")
		PrintList(got)
	})

}
