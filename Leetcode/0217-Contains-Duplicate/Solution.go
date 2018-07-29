package Solution

//	判断数组是否含有重复元素
func containsDuplicate(nums []int) bool {
	intMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		//	将nums[i]存在intMap的Key中
		if _, ok := intMap[nums[i]]; ok {
			return true
		}
		intMap[nums[i]] = i
	}
	return false
}
