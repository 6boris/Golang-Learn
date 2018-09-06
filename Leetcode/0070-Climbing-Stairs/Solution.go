package Solution

import "fmt"

//dp[i] = dp[i-1] + dp[i-2]
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	pre1 := 1
	pre2 := 2

	for i := 2; i < n; i++ {
		curr := pre1 + pre2
		pre1 = pre2
		pre2 = curr
		fmt.Println("CUR", curr, "PRE1", pre1, "PRE2", pre2)

	}
	return pre2
}
