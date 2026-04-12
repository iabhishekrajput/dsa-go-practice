package trees

import "testing"

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		tree      []int
		targetSum int
		want      bool
	}{
		{
			name:      "path exists",
			tree:      []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1},
			targetSum: 22,
			want:      true,
		},
		{
			name:      "no matching path",
			tree:      []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1},
			targetSum: 100,
			want:      false,
		},
		{
			name:      "simple tree true",
			tree:      []int{1, 2, 3},
			targetSum: 4,
			want:      true,
		},
		{
			name:      "simple tree false",
			tree:      []int{1, 2, 3},
			targetSum: 2,
			want:      false,
		},
		{
			name:      "single node match",
			tree:      []int{5},
			targetSum: 5,
			want:      true,
		},
		{
			name:      "single node no match",
			tree:      []int{5},
			targetSum: 1,
			want:      false,
		},
		{
			name:      "empty tree",
			tree:      []int{},
			targetSum: 0,
			want:      false,
		},
		{
			name:      "must be leaf not mid-node",
			tree:      []int{1, 2, -1},
			targetSum: 1,
			want:      false,
		},
		{
			name:      "negative values",
			tree:      []int{-2, -1, -3},
			targetSum: -5,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.tree)
			got := HasPathSum(root, tt.targetSum)
			if got != tt.want {
				t.Errorf("HasPathSum(target=%d) = %v, want %v", tt.targetSum, got, tt.want)
			}
		})
	}
}
