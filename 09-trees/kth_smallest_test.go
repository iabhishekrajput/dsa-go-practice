package trees

import "testing"

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		k    int
		want int
	}{
		{
			name: "1st smallest",
			tree: []int{8, 3, 10, 1, 6, -1, 14},
			k:    1,
			want: 1,
		},
		{
			name: "3rd smallest",
			tree: []int{8, 3, 10, 1, 6, -1, 14},
			k:    3,
			want: 6,
		},
		{
			name: "5th smallest",
			tree: []int{8, 3, 10, 1, 6, -1, 14},
			k:    5,
			want: 10,
		},
		{
			name: "largest element",
			tree: []int{8, 3, 10, 1, 6, -1, 14},
			k:    6,
			want: 14,
		},
		{
			name: "single node k=1",
			tree: []int{5},
			k:    1,
			want: 5,
		},
		{
			name: "left-skewed k=2",
			tree: []int{3, 2, -1, 1, -1},
			k:    2,
			want: 2,
		},
		{
			name: "simple tree k=2",
			tree: []int{2, 1, 3},
			k:    2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.tree)
			got := KthSmallest(root, tt.k)
			if got != tt.want {
				t.Errorf("KthSmallest(k=%d) = %d, want %d", tt.k, got, tt.want)
			}
		})
	}
}
