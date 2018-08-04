package Solution

import "fmt"

func Merge(l1 []int, l2 []int) []int {
	res := []int{}

	if l1 == nil {
		res = append(res, l2[:]...)
	}
	if l2 == nil {
		res = append(res, l1[:]...)
	}
	return res
}

func demo() {
	fmt.Println("demo")
}
