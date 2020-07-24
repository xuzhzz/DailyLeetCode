package dailyleetcode

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	f := make([][]int, m)
	for i := 0; i < m; i = i + 1 {
		for j := 0; j < n; j = j + 1 {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			f[i][j] = grid[i][j]
		}
	}
	// dp
	for i := m - 1; i >= 0; i = i - 1 {
		for j := n - 1; j >= 0; j = j - 1 {
			if i == m-1 && j == n-1 {
				continue
			}
			if i == m-1 {
				f[i][j] = f[i][j] + f[i][j+1]
			} else if j == n-1 {
				f[i][j] = f[i][j] + f[i+1][j]
			} else {
				f[i][j] = f[i][j] + getMin(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[0][0]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
