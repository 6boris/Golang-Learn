package main

import (
	"fmt"
	"sort"
)

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func main() {
	x := 11
	s := []int{3, 6, 8, 11, 45, 2} //注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, "在s中的位置为：", pos)
	} else {
		fmt.Println("s不包含元素", x)
	}

	GuessingGame()
}
