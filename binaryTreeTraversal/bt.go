package binarytreetraversal

// TreeNode node
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// Array2Tree 根据数组生成二叉树
func Array2Tree(nums []interface{}) *TreeNode {
	var root *TreeNode

	return root
}

// pre order
// 递归
// 非递归

// in order
// 递归
// 非递归

// post order
// 递归
// 非递归

func postOrderTraver(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	var lastTreeNode *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node == lastTreeNode {
			lastTreeNode = node
			stack = stack[0 : len(stack)-1]
			res = append(res, node.Val)
		} else {
			root = node.Right
		}
	}

	return res
}
