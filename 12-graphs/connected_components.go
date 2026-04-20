package graphs

// CountComponents returns the number of connected components in an undirected graph.
// Nodes are labeled 0 to n-1.
// edges[i] = [a, b] represents an undirected edge between a and b.
func CountComponents(n int, edges [][]int) int {
	// Your solution here

	adj := make(map[int][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	visited, components := make([]bool, n), 0
	for i := range n {
		if !visited[i] {
			components++

			queue := []int{}

			queue = append(queue, i)
			visited[i] = true

			for len(queue) > 0 {
				node := queue[0]
				queue = queue[1:]

				for _, nbr := range adj[node] {
					if !visited[nbr] {
						visited[nbr] = true
						queue = append(queue, nbr)
					}
				}

			}

		}

	}

	return components
}
