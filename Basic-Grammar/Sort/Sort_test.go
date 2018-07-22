package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("BubbleSort Sort  Test:", func(t *testing.T) {
		arr := []int{1, 3, 2, 4, 5}

		got := BubbleSort(arr)
		//want := []int{1, 3, 4, 5, 7, 9}
		want := []int{1, 2, 3, 4, 5}

		fmt.Println(got)
		fmt.Println(want)

		if !isEqual(got, want) {
			t.Error("got:", got, "want:", want)
		}
	})
}

func isEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
