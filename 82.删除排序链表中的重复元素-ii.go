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
	work := &ListNode{
		Val: -123,
	}
	work.Next = head
	flag := -124
	for work.Next != nil {
		if work.Val == work.Next.Val || work.Val == flag {
			work.Next = work.Next.Next
			flag = work.Val
		} else {
			work = work.Next
		}
	}
	return head
}

// @lc code=end
