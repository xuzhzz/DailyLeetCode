package dailyleetcode

/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// 以root为根结点的树上是否有p结点或者q结点
// 1.左右都没有 最终nil
// 2.左右都有 root
// 3.左有，右没有，最终是左
// （p=2，q=4 ）3 5 6 2return
//  (p=7, q=0) 3 5 6 2 7return 2 0return
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// @lc code=end
