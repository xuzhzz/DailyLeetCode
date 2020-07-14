package dailyleetcode

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversed := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversed
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow := head
	fast := head.Next
	work := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	headn := slow.Next
	slow.Next = nil
	headn = reverseList(headn)
	for headn != nil {
		temp := headn.Next
		headn.Next = work.Next
		work.Next = headn
		headn = temp
		work = work.Next.Next
	}
	return
}

// @lc code=end
