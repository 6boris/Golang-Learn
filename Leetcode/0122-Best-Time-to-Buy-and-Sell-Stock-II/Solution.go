package Solution

func maxProfit(prices []int) int {
	profil := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profil += prices[i] - prices[i-1]
		}
	}
	return profil
}
