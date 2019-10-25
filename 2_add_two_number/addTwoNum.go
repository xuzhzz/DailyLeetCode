package __add_two_number

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func build(val ...int) *ListNode {
	h := NewListNode()
	temp := NewListNode()
	temp = h
	for index, v := range val {
		temp.Val = v
		if index != len(val) - 1 {
			temp.Next = NewListNode()
			temp = temp.Next
		}
	}
	return h
}

func (l *ListNode) printl()  {
	temp := l
	for {

		if nil == temp {
			break
		}
		fmt.Printf("%d ==> ", temp.Val)
		temp = temp.Next
	}
	fmt.Printf("nil\n")
}

func (l *ListNode) hasNotNilNext() bool {
	if l == nil {
		return false
	}
	if l.Next == nil {
		return false
	}
	return true
}

func NewListNode() *ListNode {
	//下面的data可以用来表示链表的长度
	return &ListNode{0, nil}
}

func getVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry, sum, x, y int
		res, head *ListNode
	)
	head = NewListNode()
	res = head
	carry, sum, x, y = 0, 0, 0, 0

	for {
		if l1 == nil && l2 == nil {
			break
		}
		x = getVal(l1)
		y = getVal(l2)
		sum = x + y + carry
		carry = sum / 10
		res.Val = sum % 10
		fmt.Println(x, y, sum, carry, res.Val)
		if l1.hasNotNilNext() || l2.hasNotNilNext() || carry != 0 {
			res.Next = NewListNode()
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		res = res.Next
	}
	if carry != 0 {
		res.Val = carry
	}
	return head
}