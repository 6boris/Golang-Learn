package Solution

func countPrimes(n int) int {
	//num := n + 1
	notPrimes := [10000000]bool{}
	count := 0

	for i := 2; i < n; i++ {
		if notPrimes[i] {
			continue
		}
		count++
		// 从 i * i 开始，因为如果 k < i，那么 k * i 在之前就已经被去除过了
		for j := i * i; j < n; j = i + j {
			notPrimes[j] = true
		}
	}
	return count
}
