package dailyleetcode

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 复制节点，只复制next指针
	work := head
	for work != nil {
		newNode := &Node{
			Val: work.Val,
		}
		newNode.Next = work.Next
		work.Next = newNode
		work = work.Next.Next
	}
	// 复制random指针
	newHead := head.Next
	work = head
	for work != nil {
		if work.Random != nil {
			work.Next.Random = work.Random.Next
		}
		work = work.Next.Next
	}
	// 分离
	work = head
	temp := work.Next
	for work != nil && work.Next != nil {
		temp = work.Next
		work.Next = work.Next.Next
		work = temp
	}
	return newHead
}

// @lc code=end
