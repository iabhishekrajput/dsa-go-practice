package trees

// LowestCommonAncestor finds the lowest common ancestor of nodes p and q
// in a binary tree. Both p and q are guaranteed to exist in the tree.
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Your solution here

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
