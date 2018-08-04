package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{1, 1, 0, 1, 1, 1}
		got := findMaxConsecutiveOnes(data)
		want := 3
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

}
