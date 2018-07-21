package main

import (
	"fmt"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func TreeInit() *TreeNode {
	/*
						1
				2					3
		4			5			6			7
			10

	*/

	root := NewNode(1)
	root.left = NewNode(2)
	root.right = NewNode(3)
	root.left.left = NewNode(4)
	root.left.right = NewNode(5)
	root.right.left = NewNode(6)
	root.right.right = NewNode(7)
	root.left.left.right = NewNode(10)

	return root
}

func NewNode(val int) *TreeNode {
	n := &TreeNode{nil, nil, val}
	return n
}

//先序遍历：  节点 - 左孩子 - 右孩子
func InOrder(n *TreeNode) {
	if n == nil {
		return
	}
	fmt.Print(n.val, " ")
	InOrder(n.left)
	InOrder(n.right)
}

//后序遍历 ： 左孩子 - 右孩子 - 根结点
func MidOrder(n *TreeNode) {
	if n == nil {
		return
	}

	MidOrder(n.left)
	fmt.Print(n.val, " ")
	MidOrder(n.right)

}

//中序遍历： 左孩子  - 根结点 - 右孩子
func PostOrder(n *TreeNode) {
	if n == nil {
		return
	}
	PostOrder(n.left)
	PostOrder(n.right)
	fmt.Print(n.val, " ")
}

func main() {

	root := TreeInit()

	InOrder(root)
	fmt.Println()
	PostOrder(root)
	fmt.Println()
	MidOrder(root)
	s := lls.New()
	s.Push(1)
	s.Push(2)

	fmt.Println(s.Pop())

}
