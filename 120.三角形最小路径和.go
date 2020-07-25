package dailyleetcode

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	f := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if f[i] == nil {
				f[i] = make([]int, len(triangle[i]))
			}
			f[i][j] = triangle[i][j]
		}
	}
	// dp
	for i := len(triangle) - 2; i >= 0; i = i - 1 {
		for j := 0; j < len(triangle[i]); j = j + 1 {
			f[i][j] = getMin(f[i+1][j], f[i+1][j+1]) + triangle[i][j]
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
