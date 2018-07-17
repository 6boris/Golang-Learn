package main

import (
	"fmt"
)

//	最长不重复子串
func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwwkew"))
	fmt.Println(lengthOfLongestSubstring("pwwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("一二三二一"))
	fmt.Println(lengthOfLongestSubstring("这里是慕课网"))
	fmt.Println(lengthOfLongestSubstring("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
