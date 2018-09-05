package Solution

import (
	"testing"
)

func isEqual(data1 []int, data2 []int) bool {
	if len(data1) != len(data2) {
		return false
	}

	for i := 0; i < len(data1); i++ {
		if data1[i] != data2[i] {
			return false
		}
	}
	return true
}

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{2, 0, 2, 1, 1, 0}
		sortColors(data)
		want := []int{0, 0, 1, 1, 2, 2}
		if !isEqual(data, want) {
			t.Errorf("GOT:", data, "WANT:", want)
		}
	})
}
