package trees

// LevelOrder returns the level-order (BFS) traversal of a binary tree,
// grouped by level.
func LevelOrder(root *TreeNode) [][]int {
	// Your solution here

	result := [][]int{}

	if root == nil {
		return result
	}
	
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}
