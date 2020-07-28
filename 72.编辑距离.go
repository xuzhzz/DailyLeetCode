package dailyleetcode

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if dp[i] == nil {
				dp[i] = make([]int, n+1)
			}
			if i == 0 {
				dp[i][j] = j
			}
			if j == 0 {
				dp[i][j] = i
			}
		}
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = getMin72(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}

	}
	return dp[m][n]
}
func getMin72(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

// @lc code=end
