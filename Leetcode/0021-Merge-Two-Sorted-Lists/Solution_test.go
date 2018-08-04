package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		l1 := MakeList([]int{1, 2, 4})
		l2 := MakeList([]int{1, 3, 4})
		res := mergeTwoLists(l1, l2)
		PrintList(res)
	})

}
