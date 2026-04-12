package trees

// IsSymmetric returns true if the binary tree is a mirror of itself.
func IsSymmetric(root *TreeNode) bool {
	// Your solution here
	return checkSymmetry(root.Left, root.Right)
}

func checkSymmetry(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if (node1 == nil) || (node2 == nil) || (node1.Val != node2.Val) {
		return false
	}

	return checkSymmetry(node1.Left, node2.Right) && checkSymmetry(node1.Right, node2.Left)

}
