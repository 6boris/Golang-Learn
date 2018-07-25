package Solution

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	maxPro := 0
	minPrice := 65535
	for i := 0; i < len(prices); i++ {
		minPrice = Min(minPrice, prices[i])
		maxPro = Max(maxPro, prices[i]-minPrice)
	}
	return maxPro
}
