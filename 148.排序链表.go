package dailyleetcode

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return sortMerge(head)
}

func sortMerge(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortMerge(head)
	right := sortMerge(mid)
	// sort
	res := &ListNode{Val: 0}
	h := res
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	if left != nil {
		h.Next = left
	}
	if right != nil {
		h.Next = right
	}
	return res.Next
}

// @lc code=end
