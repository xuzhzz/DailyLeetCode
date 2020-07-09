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
	maxRes = int(math.MinInt64)
	_ = maxPathSumHelper(root)
	return maxRes
}

// 从每个节点能贡献多少出发，有可能是 左路径/右路径/或者0
func maxPathSumHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getMaxPathSum(0, maxPathSumHelper(root.Left))
	right := getMaxPathSum(0, maxPathSumHelper(root.Right))
	maxRes = getMaxPathSum(maxRes, left+right+root.Val)
	return getMaxPathSum(left+root.Val, right+root.Val, 0)
}

// @lc code=end
