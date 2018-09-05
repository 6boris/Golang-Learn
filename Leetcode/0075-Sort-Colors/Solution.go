package Solution

func sortColors(nums []int) {
	i := 0
	j := len(nums)
	for k := 0; k < j; k++ {
		//	最左边为0 满足要求，将下表右移
		if nums[k] == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			i++
			//	当前数为2，需要将其转移到末尾
		} else if nums[k] == 2 {
			j--
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}

	}
}

func sortColors2(nums []int) {
	zero := 0
	one := 0
	two := 0

	//	统计每个数值出现的次数
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero++
		}
		if nums[i] == 1 {
			one++
		}
		if nums[i] == 2 {
			two++
		}
	}
	//	根据出现的次数修改数组
	for i := 0; i < zero; i++ {
		nums[i] = 0
	}

	for i := zero; i < zero+one; i++ {
		nums[i] = 1
	}
	for i := zero + one; i < zero+one+two; i++ {
		nums[i] = 2
	}

}

func sortColors1(nums []int) {
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
