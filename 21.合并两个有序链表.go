package dailyleetcode

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	newListNode := &ListNode{Val: 0}
	work := newListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			work.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			work.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		work = work.Next
	}
	if l1 != nil {
		work.Next = l1
	}
	if l2 != nil {
		work.Next = l2
	}
	return newListNode.Next
}

// @lc code=end
