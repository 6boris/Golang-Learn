package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{2, 7, 11, 15}
		target := 9
		got := twoSum(data, target)
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
}
