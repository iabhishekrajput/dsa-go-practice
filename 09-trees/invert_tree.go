package trees

// InvertTree takes a binary tree and returns its mirror image.
// Every node's left and right children are swapped, recursively.
func InvertTree(root *TreeNode) *TreeNode {
	// Your solution here

	if root == nil {
		return nil
	}

	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}
