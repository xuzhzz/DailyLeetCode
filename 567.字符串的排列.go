package dailyleetcode

/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start

// hello\nooolleoooleh
// adc\ndcda
func checkInclusion(s1 string, s2 string) bool {
	win := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	left := 0
	right := 0
	cnt := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if need[c] != 0 {
			win[c]++
			if win[c] <= need[c] {
				cnt++
			}
		}
		for right-left >= len(s1) {
			c2 := s2[left]
			left++
			if need[c2] > 0 {
				if cnt == len(s1) {
					return true
				}
				if win[c2] <= need[c2] {
					cnt--
				}
				win[c2]--
			}
		}
	}
	return false
}

// @lc code=end
