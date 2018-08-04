package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{1, 2, 2, 4}
		got := findErrorNums(data)
		want := []int{2, 3}
		if reflect.DeepEqual(got, want) {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

}
