package backpack

// https://www.lintcode.com/problem/backpack/description
func backPack1(A []int, m int) int {
	// write your code here
	// f[i][j] 前i个物品，是否能装j
	// f[i][j] =f[i-1][j] f[i-1][j-a[i] j>a[i]
	// f[0][0]=true f[...][0]=true
	// f[n][X]
	dp := make([]int, m+1)
	for i := 0; i < len(A); i++ {
		for j := m; j > A[i]-1; j-- {
			if dp[j-A[i]]+A[i] > dp[j] {
				dp[j] = dp[j-A[i]] + A[i]
			}
		}
	}
	return dp[m]
}
