package main

import (
	"fmt"
	"math"
)

func main() {
	n, m := 3, 4
	data := []int{3, 5, 4}

	//	读取数据
	//var n int
	//var m int
	//fmt.Scanf("%d %d", &n, &m)
	//var data = make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Scanf("%d", &data[i])
	//}

	//fmt.Println(n, m, data)

	fmt.Printf("%.2f", countBound(n, m, data))
}

func countBound(n, m int, len []int) float64 {
	lbound, rbound := 1.0, math.MaxFloat32
	for lbound < rbound {
		mid := (lbound + rbound) / 2
		count := 0.0
		for i := 0; i < n; i++ {
			count += float64(len[i]) / mid
			if count < float64(m) {
				fmt.Printf("count:%.2f mid:%.2f lbound:%.2f rbound:%.2f\n", count, mid, lbound, rbound)
				rbound = mid - 1
			} else {
				fmt.Printf("11count:%.2f mid:%.2f lbound:%.2f rbound:%.2f\n", count, mid, lbound, rbound)
				lbound = mid
			}
		}
	}
	fmt.Printf("count:%.2f mid:%.2f lbound:%.2f rbound:%.2f\n", 1, 2, lbound, rbound)
	return lbound
}
