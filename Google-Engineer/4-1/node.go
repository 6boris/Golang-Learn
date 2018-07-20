package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	//	Golang 局部变量不需要知道分配在堆还是栈上
	return &treeNode{value: value}
}

//	前面括号里面的是接收者
func (node treeNode) print() {
	fmt.Println(node.value)
}

//	只有指针才能改变结构内容
func (node *treeNode) setValue(val int) {
	node.value = val
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(10)

	fmt.Println(root.left.right.value)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.print()
	root.right.left.setValue(8)
	root.right.left.print()

	pRoot := &root
	pRoot.print()
	pRoot.setValue(88)
	pRoot.print()

}
