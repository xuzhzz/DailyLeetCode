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
	nodem := head
	work := head
	//next := head.Next
	for i := 1; i < m-1; i = i + 1 {
		nodem = nodem.Next
		work = work.Next
	}
	nodem.Next = nil
	work = work.Next
	var pre, next, noden *ListNode
	for i := 0; i < n-m; i = i + 1 {
		next = work.Next
		work.Next = pre
		pre = work
		work = next
		if i == 0 {
			noden = pre
		}
	}
	nodem.Next = pre
	noden.Next = work
	return head
}

// @lc code=end
