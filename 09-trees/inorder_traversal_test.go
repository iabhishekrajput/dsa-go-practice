package trees

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		tree []int // level-order, -1 = nil
		want []int
	}{
		{
			name: "standard tree",
			tree: []int{1, 2, 3, 4, 5, -1, 6},
			want: []int{4, 2, 5, 1, 3, 6},
		},
		{
			name: "left-skewed",
			tree: []int{1, 2, -1, 3, -1},
			want: []int{3, 2, 1},
		},
		{
			name: "right-skewed",
			tree: []int{1, -1, 2, -1, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "single node",
			tree: []int{1},
			want: []int{1},
		},
		{
			name: "empty tree",
			tree: []int{},
			want: []int{},
		},
		{
			name: "complete tree",
			tree: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.tree)
			got := InorderTraversal(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
