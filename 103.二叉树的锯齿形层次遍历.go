package dailyleetcode

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	deque := make([]*TreeNode, 0)
	deque = append(deque, root)
	lv := 0
	for len(deque) > 0 {
		levelLen := len(deque)
		level := make([]int, 0)
		for i := 0; i < levelLen; i = i + 1 {
			head := deque[i]
			level = append(level, head.Val)
			if lv%2 == 0 {
			} else {
			}
		}
		deque = deque[levelLen:]
		res = append(res, level)
		lv = lv + 1
	}
	return res
}

// @lc code=end
