package dailyleetcode

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

var maxRes int = int(math.MinInt64)

func getMaxPathSum(nums ...int) int {
	max := int(math.MinInt64)
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func maxPathSum(root *TreeNode) int {
	return maxPathSumHelper(root)
}

func maxPathSumHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getMaxPathSum(0, maxPathSumHelper(root.Left)+root.Val)
	right := getMaxPathSum(0, maxPathSumHelper(root.Right)+root.Val)
	maxRes = getMaxPathSum(maxRes, maxRes+left+right)
	return getMaxPathSum(left, right)
}

// @lc code=end
