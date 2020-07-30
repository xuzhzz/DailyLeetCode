package dailyleetcode

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)
	coinsCount := len(coins)
	maxNum := 1000000
	for i := 1; i < amount+1; i++ {
		dp[i] = maxNum
		for j := 0; j < coinsCount; j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] < dp[i] {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}
	if dp[amount] == maxNum {
		return -1
	}
	return dp[amount]
}

// @lc code=end
