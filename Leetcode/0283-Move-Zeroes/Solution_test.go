package Solution

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{0, 1, 0, 3, 12}
		moveZeroes(data)
		want := []int{1, 3, 12, 0, 0}
		if !reflect.DeepEqual(data, want) {
			t.Errorf("GOT:", data, "WANT:", want)
		}
	})

}
