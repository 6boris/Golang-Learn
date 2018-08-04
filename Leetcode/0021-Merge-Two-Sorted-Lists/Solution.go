package Solution

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

//	定义结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

//	创建链表
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

//	打印链表
func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
