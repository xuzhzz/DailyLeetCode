package dailyleetcode

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// kmp算法

//  ababcdaba
//n 00120
//h 012
//t  1234

// 构建next数组
func getNext(pa string) []int {
	le := len(pa)
	next := make([]int, le)
	next[0] = 0
	head := 0
	tail := 1
	while tail < le {
		if next[head] == next[tail] {
			head = head + 1
			next[tail] = head
			tail = tail + 1
		} else {
			if head == 0 {
				next[tail] = 0
				tail = tail + 1 
			} else {
				head = next[head-1]
			}
		}
	}
	return next
}

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	// TODO KMP
	hi := 0
	ni := 0
	next := getNext(needle)
	while hi < len(haystack) {
		if haystack[hi] == needle[ni] {
			hi = hi + 1
			ni = ni + 1
		} else {
			hi = hi + next[ni]
		}
	}

}

// @lc code=end
