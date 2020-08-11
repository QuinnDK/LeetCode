package main

import "fmt"

/*
示例 1:
输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, 0)
	for i := 0; i < m; i++ {
		temp := make([]int, 0)
		for j := 0; j < n; j++ {
			temp = append(temp, 0)
		}
		dp = append(dp, temp)
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
func main() {
	nums := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(nums))
}
