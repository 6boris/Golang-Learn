package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := uniquePaths(3, 2)
		want := 3
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := uniquePaths(7, 3)
		want := 28
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
}
