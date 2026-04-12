package trees

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	// Build tree:
	//        3
	//       / \
	//      5   1
	//     / \ / \
	//    6  2 0  8
	//      / \
	//     7   4
	root := buildTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})

	// Collect node references by value for test lookups
	nodes := map[int]*TreeNode{}
	var collect func(n *TreeNode)
	collect = func(n *TreeNode) {
		if n == nil {
			return
		}
		nodes[n.Val] = n
		collect(n.Left)
		collect(n.Right)
	}
	collect(root)

	tests := []struct {
		name    string
		p, q    int
		wantVal int
	}{
		{
			name:    "split at root",
			p:       5,
			q:       1,
			wantVal: 3,
		},
		{
			name:    "ancestor is one of the nodes",
			p:       5,
			q:       4,
			wantVal: 5,
		},
		{
			name:    "deep nodes same subtree",
			p:       6,
			q:       4,
			wantVal: 5,
		},
		{
			name:    "siblings",
			p:       0,
			q:       8,
			wantVal: 1,
		},
		{
			name:    "same node",
			p:       5,
			q:       5,
			wantVal: 5,
		},
		{
			name:    "leaf and deep node",
			p:       7,
			q:       4,
			wantVal: 2,
		},
		{
			name:    "leaf and root",
			p:       6,
			q:       3,
			wantVal: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LowestCommonAncestor(root, nodes[tt.p], nodes[tt.q])
			if got == nil {
				t.Errorf("LCA(%d, %d) = nil, want %d", tt.p, tt.q, tt.wantVal)
			} else if got.Val != tt.wantVal {
				t.Errorf("LCA(%d, %d) = %d, want %d", tt.p, tt.q, got.Val, tt.wantVal)
			}
		})
	}
}
