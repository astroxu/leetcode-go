package __add_two_numbers

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := &ListNode{}
	res.Next = cur
	tmpValue := &ListNode{Val: 0}
	carry := 0
	for {
		value := l1.Val + l2.Val + carry
		cur.Next = &ListNode{Val: value % 10}
		cur = cur.Next
		carry = value / 10
		l1 = l1.Next
		l2 = l2.Next
		//精髓所在
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		if l1 == nil {
			l1 = tmpValue
		}
		if l2 == nil {
			l2 = tmpValue
		}
	}
	return res.Next.Next
}
