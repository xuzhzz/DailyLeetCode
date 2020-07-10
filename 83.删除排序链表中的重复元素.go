package dailyleetcode

/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
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
	work := head
	for work.Next != nil {
		if work.Val == work.Next.Val {
			work.Next = work.Next.Next
		} else {
			work = work.Next
		}
	}
	return head
}

// @lc code=end
