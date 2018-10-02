package main

import (
	"testing"
)

func TestPackage_1(t *testing.T) {
	t.Run("TestPackage_1_1", func(t *testing.T) {
		value := []int{2, 3, 4, 5}
		weight := []int{3, 4, 5, 6}

		got := Solution1(4, value, weight, 8)
		PrintArr(got)
		//want := 0
		//if got != want {
		//	t.Errorf("GOT:", got, " WANT:", want)
		//}

	})
}
