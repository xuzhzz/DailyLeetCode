package dailyleetcode

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	sl := len(s)
	if sl <= 1 {
		return s
	}
	p := make([][]bool, sl)
	for i := 0; i < sl; i++ {
		p[i] = make([]bool, sl)
	}
	maxL := 0
	res := ""
	for right := 0; right < sl; right++ {
		for left := 0; left <= right; left++ {
			if s[left] == s[right] && (right-left <= 2 || p[left+1][right-1]) {
				p[left][right] = true
				if right-left+1 > maxL {
					maxL = right - left + 1
					res = s[left : right+1]
				}
			}
		}
	}
	return res
}

// @lc code=end
