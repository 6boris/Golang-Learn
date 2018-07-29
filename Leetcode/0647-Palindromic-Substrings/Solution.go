package Solution

var count = 0

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		checkPalindrome(s, i, i)   //	奇数的长度
		checkPalindrome(s, i, i+1) //	偶数的长度
	}
	return count
}

func checkPalindrome(s string, start int, end int) {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		count++
		start--
		end++
	}
}
