package Solution

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func findErrorNums(nums []int) []int {
	res := make([]int, 2)

	for _, v := range nums {
		if nums[Abs(v)-1] < 0 {
			res[0] = Abs(v)
		} else {
			nums[Abs(v)-1] *= -1
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res[1] = i + 1
			break
		}
	}
	return res
}

func findErrorNums1(nums []int) []int {
	//	返回结果
	res := make([]int, 2)
	//	用户统计的Map
	m := map[int]int{}
	//	统计每个数字出现次数
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	//	找到出现2次的数，
	for k, v := range m {
		if v > 1 {
			res[0] = k
		}
	}
	//	找到消失的数
	for i := 1; i < len(nums); i++ {
		if m[i] == 0 {
			res[1] = i
			break
		}
	}
	return res
}
