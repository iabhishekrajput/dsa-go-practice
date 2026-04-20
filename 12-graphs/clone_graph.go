package graphs

// Node represents a node in an undirected graph.
type Node struct {
	Val       int
	Neighbors []*Node
}

// CloneGraph returns a deep copy of the graph starting at node.
// Every node in the returned graph is a new allocation; edges are preserved.
func CloneGraph(node *Node) *Node {
	// Your solution here

	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	return CloneGraphHelper(node, visited)
}

func CloneGraphHelper(node *Node, visited map[*Node]*Node) *Node {
	if clone, ok := visited[node]; ok {
		return clone
	}

	clone := &Node{Val: node.Val}
	visited[node] = clone

	for _, nbr := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, CloneGraphHelper(nbr, visited))
	}

	return clone
}
