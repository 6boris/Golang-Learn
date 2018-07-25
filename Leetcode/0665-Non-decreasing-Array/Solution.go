package Solution

func checkPossibility(nums []int) bool {
	cout := 0
	for i := 1; i < len(nums) && cout < 2; i++ {
		if nums[i-1] <= nums[i] {
			continue
		}
		//	不满足非递减(Non-Decreasing)序列时,即此时为递减
		if cout++; cout >= 2 {
			break
		}
		//	优先将 nums[i]降为nums[i+1]，这样可以减少 nums[i+1] > nums[i+2] 的风险。
		if i >= 2 && nums[i-2] > nums[i] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
	}
	return cout <= 1
}
