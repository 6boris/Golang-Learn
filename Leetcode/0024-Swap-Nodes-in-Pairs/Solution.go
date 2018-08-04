package Solution

import "fmt"

//	定义结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

//
func MakeList(nums []int) *ListNode {
	head := &ListNode{}
	node := head
	for i := 0; i < len(nums); i++ {
		node.Val = nums[i]
		if i != len(nums)-1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	return head
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func swapPairs(head *ListNode) *ListNode {
	node := new(ListNode)
	node.Next = head
	pre := new(ListNode)

	for pre.Next != nil && pre.Next.Next != nil {
		l1 := pre.Next
		l2 := pre.Next.Next
		next := l2.Next
		l1.Next = next
		l2.Next = l1
		pre.Next = l2
		pre = l1
	}
	return node.Next
}
