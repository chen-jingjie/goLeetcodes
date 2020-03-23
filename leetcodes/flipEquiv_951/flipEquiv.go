package flipEquiv_951

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FlipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == root2 {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val == root2.Val {
		return FlipEquiv(root1.Left, root2.Left) && FlipEquiv(root1.Right, root2.Right) ||
			(FlipEquiv(root1.Left, root2.Right) && FlipEquiv(root1.Right, root2.Left))
	}

	return false
}
