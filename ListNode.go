package dailyleetcode

import (
	"strconv"
	"strings"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func buildListNode(in string) *ListNode {
	head := &ListNode{Val: 0}
	work := head
	if in[0] == '[' && in[len(in)-1] == ']' {
		in = in[1 : len(in)-1]
	}
	for _, s := range strings.Split(in, ",") {
		s = strings.Trim(s, " ")
		newNode := &ListNode{}
		newNode.Val, _ = strconv.Atoi(s)
		work.Next = newNode
		work = work.Next
	}
	return head.Next
}

func printListNode(head *ListNode) string {
	res := ""
	for head != nil && head.Next != nil {
		res = res + strconv.Itoa(head.Val)
		res = res + "->"
		head = head.Next
	}
	res = res + strconv.Itoa(head.Val)
	return res
}

// Node random node
type NodeRandom struct {
	Val    int
	Next   *Node
	Random *Node
}

// Node Neighbors node
type Node struct {
	Val       int
	Neighbors []*Node
}
