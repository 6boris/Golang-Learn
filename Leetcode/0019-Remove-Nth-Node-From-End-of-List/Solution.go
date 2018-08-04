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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start := &ListNode{}
	fastNode := start
	slowNode := start
	slowNode.Next = head

	for n >= 0 {
		fastNode = fastNode.Next
		n--
	}

	for fastNode != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}
	slowNode.Next = slowNode.Next.Next
	return start.Next
}
