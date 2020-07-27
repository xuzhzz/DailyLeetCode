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
	for i := 0; i < len(s); i++ {
		if isPalindrome(s[0 : i+1]) {
			f[i] = 0
		} else {
			f[i] = f[i-1] + 1
			for j := 1; j < i; j++ {
				if isPalindrome(s[j : i+1]) {
					f[i] = getMin132(f[i], f[j-1]+1)
				}
			}
		}
	}
	return f[len(s)-1]
}

func isPalindrome(s string) bool {
	sl := len(s)
	for i := 0; i < sl/2; i++ {
		if s[i] != s[sl-i-1] {
			return false
		}
	}
	return true
}

func getMin132(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
