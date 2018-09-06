package Solution

func Max(data1 int, data2 int) int {
	if data1 > data2 {
		return data1
	}
	return data2
}

//dp[i] = max(dp[i-2],dp[i-3]) + nums[i]
func rob(nums []int) int {
	dp := make([][2]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = nums[i-1] + dp[i-1][0]
	}
	return Max(dp[len(nums)][0], dp[len(nums)][1])
}
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	pre3 := 0
	pre2 := 0
	pre1 := 0
	for i := 0; i < n; i++ {
		cur := Max(pre2, pre3) + nums[i]
		pre3 = pre2
		pre2 = pre1
		pre1 = cur
	}
	return Max(pre1, pre2)
}
