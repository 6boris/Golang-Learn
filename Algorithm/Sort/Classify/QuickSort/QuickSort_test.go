package QuickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 1, 2}
	QuickSort(arr)

	fmt.Println(arr)

}
