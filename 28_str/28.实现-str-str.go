package str28

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

//      a b c d a b c y
//      0 0 0 0 1 2 3 0
//
//  a a b a a a c
//  0 1 0 1 2 2 0

//  a a b a a a b a a a c

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hi := 0
	ni := 0
	next := getNext(needle)
	lh := len(haystack)
	ln := len(needle)
	for hi < lh {
		if haystack[hi] == needle[ni] {
			hi = hi + 1
			ni = ni + 1
		} else if ni != 0 {
			ni = next[ni-1]
		} else {
			hi = hi + 1
		}
		if ni == ln {
			return hi - ni
		}
		// if haystack[hi] == needle[ni] {
		// 	hi = hi + 1
		// 	ni = ni + 1
		// }
		// if ni == ln {
		// 	return hi - ni
		// } else if hi < lh && haystack[hi] != needle[ni] {
		// 	if ni != 0 {
		// 		ni = next[ni-1]
		// 	} else {
		// 		hi = hi + 1
		// 	}
		// }
	}
	return -1
}

// 构建next数组 abcy
func getNext(pa string) []int {
	le := len(pa)
	next := make([]int, le)
	next[0] = 0
	head := 0
	tail := 1
	for tail < le {
		if pa[head] == pa[tail] {
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

// @lc code=end
