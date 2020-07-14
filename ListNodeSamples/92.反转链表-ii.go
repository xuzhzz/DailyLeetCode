package dailyleetcode

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// nodem 1->nil
// pre 2->nil, work 3->4->5
// next 4->5
// pre 3->2->nil, work 4->5
// pre 4-3-2-nil work 5

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	work := &ListNode{Val: 0}
	work.Next = head
	head = work
	// 0->1->2->3->4->5->NULL
	var pre *ListNode
	for i := 0; i < m; i = i + 1 {
		pre = work
		work = work.Next
	}
	// 0->1(pre)->2(work)->3->4->5->NULL
	noden := work
	var next, temp *ListNode

	for i := 0; i < n-m+1 && work != nil; i = i + 1 {
		temp = work.Next
		work.Next = next
		next = work
		work = temp
	}
	noden.Next = work
	pre.Next = next
	return head.Next
}

// @lc code=end
