package graphs

// NumIslands counts the number of islands in a 2D grid.
// '1' is land, '0' is water. Islands are connected horizontally or vertically.
func NumIslands(grid [][]byte) int {
	// Your solution here

	if len(grid) == 0 {
		return 0
	}

	rows, columns := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, columns)
	}

	isRowColValid := func(row, col int) bool {
		return row >= 0 && row < rows && col >= 0 && col < columns && grid[row][col] == '1' && !visited[row][col]
	}

	islands := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' && !visited[row][col] {
				islands++
				queue := [][]int{}
				queue = append(queue, []int{row, col})
				visited[row][col] = true

				for len(queue) > 0 {
					node := queue[0]
					queue = queue[1:]

					var (
						topRow    = node[0] - 1
						rightRow  = node[0]
						bottomRow = node[0] + 1
						leftRow   = node[0]
						topCol    = node[1]
						rightCol  = node[1] + 1
						bottomCol = node[1]
						leftCol   = node[1] - 1
					)
					if isRowColValid(topRow, topCol) {
						visited[topRow][topCol] = true
						queue = append(queue, []int{topRow, topCol})
					}

					if isRowColValid(rightRow, rightCol) {
						visited[rightRow][rightCol] = true
						queue = append(queue, []int{rightRow, rightCol})
					}

					if isRowColValid(bottomRow, bottomCol) {
						visited[bottomRow][bottomCol] = true
						queue = append(queue, []int{bottomRow, bottomCol})
					}

					if isRowColValid(leftRow, leftCol) {
						visited[leftRow][leftCol] = true
						queue = append(queue, []int{leftRow, leftCol})
					}

				}

			}
		}
	}

	return islands
}
