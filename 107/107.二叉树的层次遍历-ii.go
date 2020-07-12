package dailyleetcode

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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

var res [][]int

func visitNextLevel(curLevel []*TreeNode) {
	nextLevel := make([]*TreeNode, 0)
	levelVal := make([]int, 0)
	if len(curLevel) > 0 {
		for _, level := range curLevel {
			levelVal = append(levelVal, level.Val)
			if level.Left != nil {
				nextLevel = append(nextLevel, level.Left)
			}
			if level.Right != nil {
				nextLevel = append(nextLevel, level.Right)
			}
		}
	}
	if len(nextLevel) > 0 {
		visitNextLevel(nextLevel)
	}
	res = append(res, levelVal)
}

func levelOrderBottom(root *TreeNode) [][]int {
	res = make([][]int, 0)
	nextLevel := make([]*TreeNode, 0)
	if root != nil {
		nextLevel = append(nextLevel, root)
		visitNextLevel(nextLevel)
	}
	return res
}

// @lc code=end
