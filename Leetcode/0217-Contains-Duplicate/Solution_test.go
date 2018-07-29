package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{1, 2, 3, 1}
		got := containsDuplicate(data)
		want := true
		if got != want {
			t.Error("GOT:", got, " WANG:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		got := containsDuplicate(data)
		want := false
		if got != want {
			t.Error("GOT:", got, " WANG:", want)
		}
	})
	t.Run("Test-3", func(t *testing.T) {
		data := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		got := containsDuplicate(data)
		want := true
		if got != want {
			t.Error("GOT:", got, " WANG:", want)
		}
	})

}
