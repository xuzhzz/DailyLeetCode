package dailyleetcode

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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

func absSub(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalance(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := isBalance(node.Left)
	right := isBalance(node.Right)
	if left == -1 || right == -1 || absSub(left, right) > 1 {
		return -1
	}
	return getMax(left, right) + 1
}

func isBalanced(root *TreeNode) bool {
	return isBalance(root) != -1
}

// @lc code=end
