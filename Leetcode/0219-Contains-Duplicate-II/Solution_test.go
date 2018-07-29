package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{1, 2, 3, 1}
		k := 3
		got := containsNearbyDuplicate(data, k)
		want := true

		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := []int{1, 0, 1, 1}
		k := 1
		got := containsNearbyDuplicate(data, k)
		want := true

		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-3", func(t *testing.T) {
		data := []int{1, 2, 3, 1, 2, 3}
		k := 2
		got := containsNearbyDuplicate(data, k)
		want := false

		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
}
