package main

import (
	"fmt"
	"strconv"
)

func main() {

	l := 0
	r := 0
	count := 0
	fmt.Scanf("%d %d", &l, &r)

	for i := l; i <= r; i++ {
		if IsEqual(i) {
			count++
		}
	}
	fmt.Println(count)

}

func IsEqual(x int) bool {
	str := strconv.Itoa(x)
	if str[0] == str[len(str)-1] {
		return true
	}

	return false
}
