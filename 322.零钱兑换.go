package dailyleetcode

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return -1
	}
	dp := make([]int, amount+1)
	coinsCount := len(coins)
	maxNum := 1000000
	for i := 1; i < amount+1; i++ {
		dp[i] = maxNum
		for j := 0; j < coinsCount; j++ {
			if i-coins[j] < 0 {
				break
			}
			dp[i] = getMin322(dp[i], dp[i-coins[j]])
		}
		if dp[i] != maxNum {
			dp[i] = dp[i] + 1
		} else {
			dp[i] = -1
		}
	}
	return dp[amount]
}

func getMin322(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
