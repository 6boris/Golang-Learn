package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{2, 3, 2}
		got := rob(data)
		want := 3
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
	t.Run("asd", func(t *testing.T) {

	})
	t.Run("Test-2", func(t *testing.T) {
		data := []int{1, 2, 3, 1}
		got := rob(data)
		want := 4
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
}
