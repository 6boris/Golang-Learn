package HeapSort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 9, 8, 7, 1}
	fmt.Println(arr)
	HeapSort(arr)
	fmt.Println(arr)
}
