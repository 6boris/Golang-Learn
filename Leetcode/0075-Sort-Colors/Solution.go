package Solution

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	left := 0
	mid := 0
	right := len(nums) - 1

	for mid <= right {
		if nums[mid] == 0 {
			nums[mid] = nums[left]
			nums[left] = 0
			mid++
			left++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid] = nums[right]
			nums[right] = 2
			right--
		}
	}
}
