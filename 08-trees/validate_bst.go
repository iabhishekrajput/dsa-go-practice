package trees

import "math"

// IsValidBST returns true if the binary tree is a valid binary search tree.
// A valid BST has: left subtree values < node < right subtree values.
func IsValidBST(root *TreeNode) bool {
	// Your solution here

	minimum, maximum := math.MinInt, math.MaxInt
	
	return validateBST(root, minimum, maximum)
}

func validateBST(root *TreeNode, minimum, maximum int) bool {
	if root == nil {
		return true
	}

	if root.Val > minimum && root.Val < maximum {
		isLeftValid := validateBST(root.Left, minimum, root.Val)
		isRightValid := validateBST(root.Right, root.Val, maximum)
	
		return isLeftValid && isRightValid
	}

	return false
}