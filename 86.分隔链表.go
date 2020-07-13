package dailyleetcode

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	work := &ListNode{Val: -999}
	work.Next = head
	tail := head
	head = work
	var temp *ListNode
	nodesCount := 1
	for tail.Next != nil {
		tail = tail.Next
		nodesCount = nodesCount + 1
	}
	for i := 0; i < nodesCount; i = i + 1 {
		if work.Next != nil && work.Next.Val >= x && work.Next.Next != nil {
			temp = work.Next
			work.Next = work.Next.Next

			temp.Next = nil
			tail.Next = temp
			tail = temp
		} else {
			work = work.Next
		}
	}
	return head.Next
}

// @lc code=end
