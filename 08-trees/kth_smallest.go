package trees

// KthSmallest returns the kth smallest element in a BST (1-indexed).
func KthSmallest(root *TreeNode, k int) int {
	// Your solution here

	var result int
	traverse(root, &k, &result)

	return result
}

func traverse(node *TreeNode, k *int, res *int) {
	if node == nil {
		return
	}

	traverse(node.Left, k, res)

	*k--
	if *k == 0 {
		*res = node.Val
		return
	}

	traverse(node.Right, k, res)
}
