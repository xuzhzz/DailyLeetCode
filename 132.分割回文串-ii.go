package dailyleetcode

/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}
	f := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if isPalindrome(s[0 : i+1]) {
			f[i] = 0
		} else {
			f[i] = f[i-1] + 1
		}
	}
	return f[len(s)-1]
}

func isPalindrome(s string) bool {
	sl := len(s)
	// 0 1 2 3 5
	for i := 0; i < sl/2; i++ {
		if s[i] != s[sl-i-1] {
			return false
		}
	}
	return true
}

// @lc code=end
