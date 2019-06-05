package main

import (
	"fmt"
)

func main() {
	var s []string
	fmt.Scanf("%s", &s)

	k := 0
	for i := 0; i < len(str); i++ {
		s[k] = s[i]
		if k >= 3 && s[k-3] == s[k-2] && s[k-2] == s[k-1] {
			k--
		}
		if k >= 4 && s[k-4] == s[k-3] && s[k-2] == s[k-1] {
			k--
		}
	}
	fmt.Println(s[:k])
}
