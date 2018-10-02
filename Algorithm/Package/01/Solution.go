package main

import "fmt"

/*
	param1:物品的数量
	param2:各个物品的重量
	param3:各个物品的价值
	param4:背包的容量
	算法说明
	m[i][j]:表示面对第i个物品，背包容量为j的时候，能获得的最大价值
	if j<weights[i]   //背包掏空了都放不下第i个物品
		m[i][j]=m[i-1][j]
	else   //背包掏空了可以放下第i个物品
		m[i][j]=max(m[i-1][j],m[i-1][j-w[i]]+v[i])
*/
func Solution1(n int, weights []int, values []int, c int) [][]int {
	//定义二维数组
	m := make([][]int, (n + 1))
	for i := 0; i < n+1; i++ {
		m[i] = make([]int, (c + 1))
	}
	//初始化m[1][1...c]
	for i := 1; i <= c; i++ {
		if weights[0] > i {
			m[1][i] = 0
		} else {
			m[1][i] = values[0]
		}
	}
	//求数组中剩余变量的值
	for i := 2; i <= n; i++ {
		for j := 1; j <= c; j++ {
			if j < weights[i-1] {
				m[i][j] = m[i-1][j]
			} else {
				//fmt.Println(i,j,j-weights[i-1])
				m[i][j] = Max(m[i-1][j], m[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}
	return m
}

func PrintArr(x [][]int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			fmt.Print(" ", x[i][j])
		}
		fmt.Println()
	}
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
