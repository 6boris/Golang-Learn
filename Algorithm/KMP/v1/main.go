package main

import (
	"fmt"
)

func GetNextValueArray(sub []byte) (next []int) {
	var (
		length        int = len(sub)
		middle        int
		compare_left  int
		compare_right int
		match_count   int
	)

	next = make([]int, length)
	next[0] = 0
	next[1] = 0

	for i := 2; i < length; i++ {
		middle = i / 2
		match_count = 0

		if i%2 == 0 {
			for j := 0; j < middle; j++ {
				compare_left = 0
				compare_right = i - 1 - j
				for compare_left <= j {
					if sub[compare_left] != sub[compare_right] {
						break
					}
					compare_left++
					compare_right++
				}
				if compare_left == j+1 {
					match_count++
				}
			}
			next[i] = match_count

		} else {
			for j := 0; j <= middle; j++ {
				compare_left = 0
				compare_right = i - 1 - j
				for compare_left <= j {
					if sub[compare_left] != sub[compare_right] {
						break
					}
					compare_left++
					compare_right++
				}
				if compare_left == j+1 {
					match_count++
				}
			}
			next[i] = match_count
		}
	}

	return next
}

func ReviseNextValueArray(next []int) []int {
	var length int = len(next)
	for i := 2; i < length; i++ {
		if next[i] == next[next[i]] {
			next[i] = next[next[i]]
		}
	}

	return next
}

//在content中的start-end之间寻找sub子串
//成功返回匹配成功的起始下标，匹配失败则返回-1
func KMP(content []byte, start_index int, end_index int, sub []byte) (index int) {
	var (
		next       []int = ReviseNextValueArray(GetNextValueArray(sub))
		sub_index  int   = 0
		sub_length int   = len(sub)
	)
	for i := start_index; i <= end_index; i++ {
		if content[i] == sub[sub_index] {
			match_start := i
			for j := sub_index; j <= sub_length; j++ {
				if j == sub_length {
					return match_start - sub_index
				}
				if i >= end_index || content[i] != sub[j] {
					sub_index = next[j]
					break
				}
				i++
			}
		}
	}

	return -1
}

func main() {
	content := []byte("why every programming language use the hello world as the first test???")
	sub := []byte("hello world")
	fmt.Println(KMP(content, 0, len(content)-1, sub))
}
