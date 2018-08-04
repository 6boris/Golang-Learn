package Solution

//	双指针
//	nums[n+1], index is [0, n], num is [1, n]
//	consider as cycle linked list, nums[i] is the index of next position.
// 	No way to jump back to index = 0, so index = 0 is a start point
//	use slow/fast pointers
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	for slow == 0 || slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

//	Hash Map
func findDuplicate2(nums []int) int {
	found := -1
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > 1 {
			found = v
			break
		}
	}
	return found
}

//	二分查找
func findDuplicate1(nums []int) int {
	low := 1
	high := len(nums) - 1

	for low < high {
		mid := low + (high-low)/2
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count <= mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
