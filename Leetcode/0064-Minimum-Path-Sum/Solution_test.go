package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}
		got := minPathSum(data)
		want := 7
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
}
