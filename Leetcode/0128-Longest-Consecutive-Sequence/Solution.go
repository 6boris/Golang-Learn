package Solution

func longestConsecutive(nums []int) int {
	intMap := map[int]bool{}

	//	Build the map, and int this way we also remove duplicates, takes O(n) to create the map
	for _, v := range nums {
		intMap[v] = true
	}
	maxL := 1
	toLeft, toRight := true, true
	for k := range intMap {
		curL := 1
		for i := 1; intMap[k-i] || intMap[k+i]; i++ {
			toLeft = toLeft && intMap[k-i]
			toRight = toRight && intMap[k+i]
			if toLeft {
				curL++
				delete(intMap, k-i)
			}
			if toRight {
				curL++
				delete(intMap, k+i)
			}
		}
		if curL > maxL {
			maxL = curL
		}
	}
	return maxL

}
