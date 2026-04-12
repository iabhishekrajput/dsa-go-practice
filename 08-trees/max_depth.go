package trees

// MaxDepth returns the maximum depth of a binary tree.
// Depth is the number of nodes along the longest root-to-leaf path.
func MaxDepth(root *TreeNode) int {
	// Your solution here

	if root == nil {
		return 0
	}

	leftMaxDepth := MaxDepth(root.Left)
	rightMaxDepth := MaxDepth(root.Right)

	return 1 + max(leftMaxDepth, rightMaxDepth)
}
