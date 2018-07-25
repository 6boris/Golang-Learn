package main

import "fmt"

//func isSubsequence(s string, t string) bool {
//	last := 0
//	for i := 0; i < len(s); i++ {
//		for j := last; j < len(t); j++ {
//			if s[i] == t[j] {
//				last = j + 1
//				break
//			}
//			if j == len(t) {
//				return false
//			}
//		}
//	}
//	return true
//}

//	暴力遍历
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
