package dailyleetcode

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
// [4,10,4,3,8,9]
// [10,9,2,5,3,7,101,18]
func lengthOfLIS(nums []int) int {
	lr := len(nums)
	if lr <= 1 {
		return lr
	}
	dp := make([]int, lr) // 以dp[i]结尾的序列最长上升长度
	dp[0] = 1
	res := 0
	for i := 1; i < lr; i++ {
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > nums[i] {
				j--
			}
		}
		if j >= 0 {
			dp[i] = dp[j] + 1
		} else {
			dp[i] = 1
		}
		res = getMax300(res, dp[i])
	}
	return res
}
func getMax300(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
