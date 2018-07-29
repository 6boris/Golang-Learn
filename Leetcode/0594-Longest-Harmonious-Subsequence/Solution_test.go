package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{1, 3, 2, 2, 5, 2, 3, 7}
		got := findLHS(data)
		want := 5
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
	t.Run("Test-2", func(t *testing.T) {
		data := []int{1, 3, 2, 2, 5, 1, 3, 7}
		got := findLHS(data)
		want := 4
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
}
