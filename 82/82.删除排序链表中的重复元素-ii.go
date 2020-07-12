package dailyleetcode

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//  dummy node ?
	work := &ListNode{Val: -1}
	work.Next = head
	head = work
	for work.Next != nil && work.Next.Next != nil {
		if work.Next.Val == work.Next.Next.Val {
			dupNum := work.Next.Val
			for work.Next != nil && work.Next.Val == dupNum {
				work.Next = work.Next.Next
			}
		} else {
			work = work.Next
		}
	}
	return head.Next
}

// @lc code=end
