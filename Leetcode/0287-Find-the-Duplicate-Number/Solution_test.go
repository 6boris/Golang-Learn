package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{1, 3, 4, 2, 2}
		got := findDuplicate(data)
		want := 2
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
	t.Run("Test-2", func(t *testing.T) {
		data := []int{3, 1, 3, 4, 2}
		got := findDuplicate(data)
		want := 3
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

}
