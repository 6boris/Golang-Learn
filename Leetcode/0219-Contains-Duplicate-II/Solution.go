package Solution

//	判断最大差值
func containsNearbyDuplicate(nums []int, k int) bool {
	intMap := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := intMap[nums[i]]; ok {
			return true
		}
		if i > k-1 {
			delete(intMap, nums[i-k])
		}

		intMap[nums[i]] = true
	}
	return false
}
