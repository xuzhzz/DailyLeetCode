package dailyleetcode

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func traversal(node *TreeNode) int {
	if node == nil {
		return 0
	}
	ldep := traversal(node.Left)
	rdep := traversal(node.Right)
	if ldep > rdep {
		return 1 + ldep
	}
	return 1 + rdep
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return traversal(root)
}

// @lc code=end
