package trees

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want bool
	}{
		{
			name: "symmetric tree",
			tree: []int{1, 2, 2, 3, 4, 4, 3},
			want: true,
		},
		{
			name: "asymmetric values",
			tree: []int{1, 2, 2, -1, 3, -1, 3},
			want: false,
		},
		{
			name: "single node",
			tree: []int{1},
			want: true,
		},
		{
			name: "two levels symmetric",
			tree: []int{1, 2, 2},
			want: true,
		},
		{
			name: "two levels asymmetric",
			tree: []int{1, 2, 3},
			want: false,
		},
		{
			name: "asymmetric structure",
			tree: []int{1, 2, 2, 3, -1, -1, -1},
			want: false,
		},
		{
			name: "deep symmetric",
			tree: []int{1, 2, 2, 3, 4, 4, 3, 5, -1, -1, 6, 6, -1, -1, 5},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.tree)
			got := IsSymmetric(root)
			if got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
