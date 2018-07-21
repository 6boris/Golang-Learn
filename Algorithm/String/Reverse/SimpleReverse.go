package main

import (
	"fmt"
)

func SimpleReverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	SimpleReverseInt(s[:2])
	SimpleReverseInt(s[2:])
	SimpleReverseInt(s)
	fmt.Println(s)

	Reverse("a")

}

func Reverse(str string) bool {
	fmt.Println(str)
	return false
}
