package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		root := new(TreeNode)

		root.Val = 3
		root.Left = new(TreeNode)
		root.Left.Val = 9
		root.Right = new(TreeNode)
		root.Right.Val = 20
		root.Left.Left = nil
		root.Left.Right = nil
		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 15
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 7
		root.Right.Left.Left = nil
		root.Right.Left.Right = nil
		root.Right.Right.Left = nil
		root.Right.Right.Right = nil

		got := maxDepth(root)
		want := 3
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
		append()
	})
}
