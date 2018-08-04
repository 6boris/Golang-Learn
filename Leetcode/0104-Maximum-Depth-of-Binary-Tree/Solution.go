package Solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func buildTree(nums []int) *TreeNode {
	root := new(TreeNode)

	for i := 0; i < len(nums); i++ {

	}
}

func build() {

}
