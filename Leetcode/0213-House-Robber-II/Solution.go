package Solution

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return Max(getRob(nums, 0, n-2), getRob(nums, 1, n-1))

}

func getRob(nums []int, frist int, last int) int {
	pre1, pre2, pre3 := 0, 0, 0

	for i := frist; i <= last; i++ {
		cur := Max(pre3, pre2) + nums[i]
		pre3 = pre2
		pre2 = pre1
		pre1 = cur
	}
	return Max(pre1, pre2)
}
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
