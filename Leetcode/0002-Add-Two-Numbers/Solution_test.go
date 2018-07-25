package Solution

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestaddTwoNumbers(t testing.T) {
	l1 := ListNode{}
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3

	l2 := ListNode{}
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4
	fmt.Println(l1)
}
