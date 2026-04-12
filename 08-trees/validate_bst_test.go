package trees

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want bool
	}{
		{
			name: "valid BST",
			tree: []int{2, 1, 3},
			want: true,
		},
		{
			name: "invalid right child",
			tree: []int{5, 1, 4, -1, -1, 3, 6},
			want: false,
		},
		{
			name: "single node",
			tree: []int{1},
			want: true,
		},
		{
			name: "valid larger BST",
			tree: []int{8, 3, 10, 1, 6, -1, 14},
			want: true,
		},
		{
			name: "invalid grandchild",
			tree: []int{5, 1, 6, -1, -1, 3, 7},
			want: false,
		},
		{
			name: "equal values not valid",
			tree: []int{2, 2, 3},
			want: false,
		},
		{
			name: "left-skewed valid",
			tree: []int{3, 2, -1, 1, -1},
			want: true,
		},
		{
			name: "right-skewed valid",
			tree: []int{1, -1, 2, -1, 3},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.tree)
			got := IsValidBST(root)
			if got != tt.want {
				t.Errorf("IsValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
