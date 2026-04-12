package trees

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want int
	}{
		{
			name: "standard tree",
			tree: []int{1, 2, 3, 4, 5, -1, 6},
			want: 3,
		},
		{
			name: "single node",
			tree: []int{1},
			want: 1,
		},
		{
			name: "empty tree",
			tree: []int{},
			want: 0,
		},
		{
			name: "left-skewed",
			tree: []int{1, 2, -1, 3, -1},
			want: 3,
		},
		{
			name: "right-skewed",
			tree: []int{1, -1, 2, -1, 3},
			want: 3,
		},
		{
			name: "balanced two levels",
			tree: []int{1, 2, 3},
			want: 2,
		},
		{
			name: "asymmetric deeper left",
			tree: []int{3, 9, 20, -1, -1, 15, 7},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.tree)
			got := MaxDepth(root)
			if got != tt.want {
				t.Errorf("MaxDepth() = %d, want %d", got, tt.want)
			}
		})
	}
}
