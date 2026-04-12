package trees

// SearchBST finds a node with the given value in a BST.
// Returns the node (with its subtree), or nil if not found.
func SearchBST(root *TreeNode, val int) *TreeNode {
	// Your solution here

	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		return SearchBST(root.Right, val)
	}

	return SearchBST(root.Left, val)
}
