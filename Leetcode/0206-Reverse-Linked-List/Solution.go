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

func deleteDuplicates(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	} else {
		return head
	}
}
