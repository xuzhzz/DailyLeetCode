package dailyleetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start

func decodeString(s string) string {
	if len(s) == 4 {
		return s
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i = i + 1 {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			// get [ index
			idx := i
			for idx >= 0 && s[idx] != '[' {
				idx = idx - 1
			}
			// get str
			repeatedS := s[idx+1 : i]

			//  get repeat times
			repeatedTimes := 0
			idx = idx - 1
			for idx >= 0 && s[idx] >= '0' && s[idx] <= '9' {
				num, _ := strconv.Atoi(string(s[idx]))
				repeatedTimes = repeatedTimes*10 + num
				idx = idx - 1
			}
			stack = stack[:idx+1]
			rps := ""
			for repeatedTimes > 0 {
				rps = rps + repeatedS
				repeatedTimes = repeatedTimes - 1
			}
			for idx = 0; idx < len(rps); idx = idx + 1 {
				stack = append(stack, rps[idx])
			}
		}
	}
	return string(stack)
}

// @lc code=end
