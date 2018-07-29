package Solution

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLHS(nums []int) int {
	intMap := map[int]int{}

	//	Use a HashMap to save every element and it's occurrences
	//	nums[1 3 2 2 5 2 3 7]
	//	map[1:1 3:2 2:3 5:1 7:1]
	for _, v := range nums {
		intMap[v]++
	}
	max := 0
	//	Every Each will Calculate the max element from front middle behind element
	//	Don't  worry the out of the range ,because of it use map to store

	for k, v := range intMap {
		var v1, v2 int
		//	If this front element of this element has ever appeared
		if intMap[k-1] > 0 {
			v1 = intMap[k+1] + v
		}
		//	If this behind element of this element has ever appeared
		if intMap[k+1] > 0 {
			v1 = intMap[k+1] + v
		}
		//	Get the max element  from front middle behind element
		max = Max(Max(v1, v2), max)
	}

	return max
}
