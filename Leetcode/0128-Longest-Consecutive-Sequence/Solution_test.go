package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{100, 4, 200, 1, 3, 2}
		got := longestConsecutive(data)
		want := 4
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

}
