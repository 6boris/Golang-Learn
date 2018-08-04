package Solution

//	Golang的切片是一个非常特殊的类型，
//	此时的nums是介于传值和传引用之间的
//	类型,值会改变，但是len 和cap 并不会改变
//	所以才会有nums = append(nums, 1)这种骚操作
func moveZeroes(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	//	将不是0的元素依次在前面排列，最后把其他的补0
	insertPos := 0

	for _, v := range nums {
		if v != 0 {
			nums[insertPos] = v
			insertPos++
		}
	}
	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}
