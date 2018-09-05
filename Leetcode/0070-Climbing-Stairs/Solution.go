package Solution

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	pre1 := 2
	pre2 := 1

	for i := 2; i < n; i++ {
		cur := pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}
