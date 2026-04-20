package graphs

import (
	"sort"
	"testing"
)

// buildGraph constructs a graph from an adjacency-list representation.
// adj[i] = list of neighbor values for node with value i+1.
// Returns node with value 1 (or nil if adj is empty).
func buildGraph(adj [][]int) *Node {
	if len(adj) == 0 {
		return nil
	}
	nodes := make([]*Node, len(adj))
	for i := range adj {
		nodes[i] = &Node{Val: i + 1}
	}
	for i, neighbors := range adj {
		for _, n := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[n-1])
		}
	}
	return nodes[0]
}

// collectGraph does a BFS and returns: visited values, edges (sorted pairs), node pointers keyed by value.
func collectGraph(start *Node) (values []int, edges [][2]int, ptrs map[int]*Node) {
	ptrs = make(map[int]*Node)
	if start == nil {
		return
	}
	visited := make(map[*Node]bool)
	queue := []*Node{start}
	visited[start] = true
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		ptrs[n.Val] = n
		values = append(values, n.Val)
		for _, nb := range n.Neighbors {
			a, b := n.Val, nb.Val
			if a > b {
				a, b = b, a
			}
			edges = append(edges, [2]int{a, b})
			if !visited[nb] {
				visited[nb] = true
				queue = append(queue, nb)
			}
		}
	}
	sort.Ints(values)
	// Dedupe edges (each undirected edge shows up twice)
	edgeSet := make(map[[2]int]bool)
	for _, e := range edges {
		edgeSet[e] = true
	}
	edges = edges[:0]
	for e := range edgeSet {
		edges = append(edges, e)
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][0] != edges[j][0] {
			return edges[i][0] < edges[j][0]
		}
		return edges[i][1] < edges[j][1]
	})
	return
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name string
		adj  [][]int
	}{
		{"square graph", [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}},
		{"single node no neighbors", [][]int{{}}},
		{"two nodes connected", [][]int{{2}, {1}}},
		{"triangle", [][]int{{2, 3}, {1, 3}, {1, 2}}},
		{"star graph", [][]int{{2, 3, 4, 5}, {1}, {1}, {1}, {1}}},
		{"line graph", [][]int{{2}, {1, 3}, {2, 4}, {3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adj)
			clone := CloneGraph(original)

			origVals, origEdges, origPtrs := collectGraph(original)
			cloneVals, cloneEdges, clonePtrs := collectGraph(clone)

			// Structure must match
			if len(origVals) != len(cloneVals) {
				t.Fatalf("node count mismatch: original=%d clone=%d", len(origVals), len(cloneVals))
			}
			for i := range origVals {
				if origVals[i] != cloneVals[i] {
					t.Errorf("value mismatch at index %d: original=%d clone=%d", i, origVals[i], cloneVals[i])
				}
			}
			if len(origEdges) != len(cloneEdges) {
				t.Fatalf("edge count mismatch: original=%d clone=%d", len(origEdges), len(cloneEdges))
			}
			for i := range origEdges {
				if origEdges[i] != cloneEdges[i] {
					t.Errorf("edge mismatch at index %d: original=%v clone=%v", i, origEdges[i], cloneEdges[i])
				}
			}

			// Deep copy: every node pointer in clone must be distinct from original
			for val, op := range origPtrs {
				cp, ok := clonePtrs[val]
				if !ok {
					t.Errorf("clone missing node with val=%d", val)
					continue
				}
				if op == cp {
					t.Errorf("node val=%d was not cloned (same pointer)", val)
				}
			}
		})
	}

	t.Run("nil input", func(t *testing.T) {
		if got := CloneGraph(nil); got != nil {
			t.Errorf("CloneGraph(nil) = %v, want nil", got)
		}
	})
}
