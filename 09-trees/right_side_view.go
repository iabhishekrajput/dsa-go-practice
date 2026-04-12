package trees

// RightSideView returns the values of nodes visible from the right side
// of the binary tree, from top to bottom.
func RightSideView(root *TreeNode) []int {
	// Your solution here

	result := []int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentlyVisible := 0

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			currentlyVisible = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentlyVisible)
	}

	return result
}
