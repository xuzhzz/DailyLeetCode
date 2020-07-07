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
func treeDepMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ldep := treeDepMax(root.Left)
	rdep := treeDepMax(root.Right)
	if ldep > rdep {
		return 1 + ldep
	}
	return 1 + rdep
}
func treeDepMin(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ldep := treeDepMax(root.Left)
	rdep := treeDepMax(root.Right)
	if ldep < rdep {
		return 1 + ldep
	}
	return 1 + rdep
}

func isBalanced(root *TreeNode) bool {
	// if root == nil || (root.Left == nil && root.Right == nil) {
	// 	return false
	// }
	maxDep := treeDepMax(root)
	minDep := treeDepMin(root)
	return maxDep-minDep <= 1
}

// @lc code=end
