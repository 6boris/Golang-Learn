package Solution

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestListNodeQueue_Pop(t *testing.T) {
	h := ListNodeQueue{}
	heap.Init(&h)
	h.Push(&ListNode{Val: 1})
	h.Push(&ListNode{Val: 2})

	res := heap.Pop(&h).(*ListNode)
	fmt.Println(res.Val)
	fmt.Println(h.Pop().(*ListNode).Val)
}
