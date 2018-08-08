package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {

	t.Run("Test_Solution_Merge", func(t *testing.T) {
		l1 := []int{1, 2, 4}
		l2 := []int{1, 3, 4}
		fmt.Println(l1, l2)
	})

	t.Run("Test_Demo", func(t *testing.T) {
		demo()
	})
}
