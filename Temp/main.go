package main

import "fmt"

func main() {
	fmt.Println(GetPage(1, 6, 53))
}

func GetPage(left, mid, right int) [7]int {
	arr := [7]int{}
	MaxWidth := 3

	i, j, count := mid+1, 4, 0
	for i <= right && count < MaxWidth {
		arr[j] = i
		i++
		count++
	}

	j, j = 0, mid-1
	for i >= left && j < MaxWidth {

		i--
		j++
	}

	return arr
}
