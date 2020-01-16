package __add_two_numbers

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	var (
		l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
		l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	)
	forPrint(l1)
	forPrint(l2)

	res := addTwoNumbers(l1, l2)
	forPrint(res)
}

func forPrint(l *ListNode) {
	if l == nil {
		return
	}
	for {
		fmt.Printf("%v", l.Val)
		if l.Next == nil {
			fmt.Println()
			break
		}
		fmt.Print("->")
		l = l.Next
	}
}
