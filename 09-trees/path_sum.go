package trees

// HasPathSum returns true if the tree has a root-to-leaf path
// that adds up to the target sum.
func HasPathSum(root *TreeNode, targetSum int) bool {
	// Your solution here
	return traverseHasPathSum(root, targetSum)
}

func traverseHasPathSum(node *TreeNode, remianing int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return node.Val == remianing
	}

	return traverseHasPathSum(node.Left, remianing-node.Val) || traverseHasPathSum(node.Right, remianing-node.Val)
}
