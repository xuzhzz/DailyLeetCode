/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	f := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		f[i] = false
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if canGetBridge(f, i, nums[i]) {
			f[i] = true
		}
	}
	return f[0]
}

func canGetBridge(f []bool, i, jumps int) bool {
	if i+jumps >= len(f)-1 {
		return true
	}
	for idx := 1; idx <= jumps; idx++ {
		if f[i+idx] {
			return true
		}
	}
	return false
}

// @lc code=end

