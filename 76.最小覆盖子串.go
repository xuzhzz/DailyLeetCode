package dailyleetcode

import "math"

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	// if len(s) == 0 || len(t) == 0 {
	// 	return ""
	// }
	// win := make(map[byte]int)
	// tMap := make(map[byte]int)
	// for i := 0; i < len(t); i++ {
	// 	tMap[t[i]]++
	// }
	// left := 0
	// right := 0
	// start := 0
	// end := 0
	// matchChars := 0
	// minL := 10000000
	// for right < len(s) {
	// 	c := s[right]
	// 	right++
	// 	if tMap[c] > 0 {
	// 		win[c]++
	// 		if win[c] < tMap[c] {
	// 			matchChars++
	// 		}
	// 	}
	// 	for matchChars == len(t) {
	// 		if right-left < minL {
	// 			minL = right - left
	// 			start = left
	// 			end = right + 1
	// 		}
	// 		sc := s[left]
	// 		left++
	// 		if tMap[sc] > 0 {
	// 			if win[sc] == tMap[sc] {
	// 				matchChars--
	// 			}
	// 			win[sc]--
	// 		}
	// 	}
	// }
	// return s[start:end]
	// 保存滑动窗口字符集
	win := make(map[byte]int)
	// 保存需要的字符集
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	// 窗口
	left := 0
	right := 0
	// match匹配次数
	match := 0
	start := 0
	end := 0
	min := math.MaxInt64
	var c byte
	for right < len(s) {
		c = s[right]
		right++
		// 在需要的字符集里面，添加到窗口字符集里面
		if need[c] != 0 {
			win[c]++
			// 如果当前字符的数量匹配需要的字符的数量，则match值+1
			if win[c] == need[c] {
				match++
			}
		}

		// 当所有字符数量都匹配之后，开始缩紧窗口
		for match == len(need) {
			// 获取结果
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c = s[left]
			left++
			// 左指针指向不在需要字符集则直接跳过
			if need[c] != 0 {
				// 左指针指向字符数量和需要的字符相等时，右移之后match值就不匹配则减一
				// 因为win里面的字符数可能比较多，如有10个A，但需要的字符数量可能为3
				// 所以在压死骆驼的最后一根稻草时，match才减一，这时候才跳出循环
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}

// @lc code=end
