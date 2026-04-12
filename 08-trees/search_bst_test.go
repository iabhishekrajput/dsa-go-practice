package trees

import "testing"

func TestSearchBST(t *testing.T) {
	// Build BST: [8, 3, 10, 1, 6, nil, 14]
	tree := []int{8, 3, 10, 1, 6, -1, 14}

	tests := []struct {
		name    string
		val     int
		wantVal int
		wantNil bool
	}{
		{
			name:    "find root",
			val:     8,
			wantVal: 8,
		},
		{
			name:    "find left child",
			val:     3,
			wantVal: 3,
		},
		{
			name:    "find right child",
			val:     10,
			wantVal: 10,
		},
		{
			name:    "find leaf",
			val:     1,
			wantVal: 1,
		},
		{
			name:    "find deep node",
			val:     6,
			wantVal: 6,
		},
		{
			name:    "not found",
			val:     5,
			wantNil: true,
		},
		{
			name:    "not found larger",
			val:     20,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tree)
			got := SearchBST(root, tt.val)
			if tt.wantNil {
				if got != nil {
					t.Errorf("SearchBST(%d) = %d, want nil", tt.val, got.Val)
				}
				return
			}
			if got == nil {
				t.Errorf("SearchBST(%d) = nil, want %d", tt.val, tt.wantVal)
			} else if got.Val != tt.wantVal {
				t.Errorf("SearchBST(%d).Val = %d, want %d", tt.val, got.Val, tt.wantVal)
			}
		})
	}

	t.Run("empty tree", func(t *testing.T) {
		got := SearchBST(nil, 1)
		if got != nil {
			t.Errorf("SearchBST on nil tree = %d, want nil", got.Val)
		}
	})
}
