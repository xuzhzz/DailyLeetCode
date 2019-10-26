package longest_substring_without_repeating_characters

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// abcabcbb
// left max idx hash[a]  hash[b]   hash[c]
// 0   1   0     1
// 0   2   1     1       2
//
func LengthOfLongestSubstring(s string) int {
	var (
		leftIdx int
		maxLen int
		hashMap map[rune] int
	)
	leftIdx = 0
	maxLen = 1
	hashMap = make(map[rune] int)
	for idx, v := range s {
		if val, ok := hashMap[v]; ok {
			leftIdx = getMax(val, leftIdx)
			hashMap[v] = idx+1
		} else {
			hashMap[v] = idx+1
		}
		maxLen = getMax(maxLen, idx - leftIdx + 1)
	}

	return maxLen
}
