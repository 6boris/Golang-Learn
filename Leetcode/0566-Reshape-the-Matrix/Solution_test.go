package Solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		//nums := [][]int{[[1,2], [3,4]]}
		nums := [][]int{{1, 2}, {3, 4}}
		got := matrixReshape(nums, 1, 4)
		want := [][]int{{1, 2, 3, 4}}
		if reflect.DeepEqual(got, want) {
			t.Errorf("GOT:", got, "WANT:", want)
		}

		fmt.Println(0 % 4)
		fmt.Println(1 % 4)
		fmt.Println(2 % 4)
		fmt.Println(3 % 4)
	})
}
