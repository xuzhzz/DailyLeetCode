package dailyleetcode

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 && len(obstacleGrid[0]) == 0 {
		return 0
	} else if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			if obstacleGrid[i][j] == 1 {
				f[i][j] = -1
			} else if i == 0 || j == 0 {
				if (i > 0 && f[i-1][j] == -1) || (j > 0 && f[i][j-1] == -1) {
					f[i][j] = -1
				} else {
					f[i][j] = 1
				}
			}
		}
	}
	// dp
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if f[i][j] != -1 {
				if f[i-1][j] == -1 && f[i][j-1] == -1 {
					f[i][j] = -1
				} else if f[i-1][j] == -1 {
					f[i][j] = f[i][j-1]
				} else if f[i][j-1] == -1 {
					f[i][j] = f[i-1][j]
				} else {
					f[i][j] = f[i][j-1] + f[i-1][j]
				}
			}
		}
	}
	if f[m-1][n-1] != -1 {
		return f[m-1][n-1]
	}
	return 0
}

// @lc code=end
