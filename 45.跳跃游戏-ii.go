package dailyleetcode

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
// f 上每一点表示 以这个点为起点跳到最后一点需要的最少步
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	endPos := len(nums) - 1
	f := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= endPos {
			f[i] = 1
		} else {
			minStep := 100000000
			for j := i; j <= i+nums[i]; j++ {
				if minStep > f[j] && f[j] != 0 {
					minStep = f[j]
				}
			}
			if minStep != 100000000 {
				f[i] = minStep + 1
			}
		}
	}
	return f[0]
}

// @lc code=end
