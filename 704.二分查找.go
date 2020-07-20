package dailyleetcode

/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	for len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}
	r := len(nums) - 1
	l := 0
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// @lc code=end
