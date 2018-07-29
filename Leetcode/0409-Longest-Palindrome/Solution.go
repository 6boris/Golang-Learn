package Solution

func longestPalindrome(s string) int {
	num := len(s)
	if num <= 1 {
		return num
	}
	//	将每个字符存到Map中
	strMap := map[rune]int{}
	for _, v := range s {
		strMap[v]++
	}

	//	遍历Hash表，判断字符出现的次数如果为奇数则num-1,如果最后出现过奇数数量+1
	//
	HasOdd := false
	for _, v := range strMap {
		if v%2 != 0 {
			num -= 1
			HasOdd = true
		}
	}
	if HasOdd {
		num += 1
	}
	return num

}

//
//1.如果为双不动
//2.如果为单 数量减一
//3.如果有单，最后结果减一
