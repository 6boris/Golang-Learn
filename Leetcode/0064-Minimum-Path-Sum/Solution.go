package Solution

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	//m, n := len(grid), len(grid[0])
	m := len(grid)
	n := len(grid[0])

	//dp := [m][n]int{{}}
	var dp [3][3]int

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = grid[i][j] + Min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

func Min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
