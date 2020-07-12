package dailyleetcode

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func dfsRight(root *TreeNode) int {
	if root.Right != nil {
		return dfsRight(root.Right)
	}
	return root.Val
}
func dfsLeft(root *TreeNode) int {
	if root.Left != nil {
		return dfsLeft(root.Left)
	}
	return root.Val
}
func validBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && dfsRight(root.Left) >= root.Val {
		return false
	}
	if root.Right != nil && dfsLeft(root.Right) <= root.Val {
		return false
	}
	return validBST(root.Left) && validBST(root.Right)
}

func isValidBST(root *TreeNode) bool {
	return validBST(root)
}

// @lc code=end
