package Solution

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	cur := 0
	for _, v := range nums {
		if v == 0 {
			cur = 0
		} else {
			cur += 1
		}
		max = Max(max, cur)
	}
	return max
}
