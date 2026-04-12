package trees

// InorderTraversal returns the inorder (Left → Root → Right) traversal of a binary tree.
func InorderTraversal(root *TreeNode) []int {
	// Your solution here

	if root == nil {
		return []int{}
	}

	left := InorderTraversal(root.Left)
	right := InorderTraversal(root.Right)

	result := append(left, root.Val)
	result = append(result, right...)

	return result
}
