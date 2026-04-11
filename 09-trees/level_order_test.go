package trees

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want [][]int
	}{
		{
			name: "standard tree",
			tree: []int{1, 2, 3, 4, 5, -1, 6},
			want: [][]int{{1}, {2, 3}, {4, 5, 6}},
		},
		{
			name: "single node",
			tree: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "empty tree",
			tree: []int{},
			want: [][]int{},
		},
		{
			name: "left-skewed",
			tree: []int{1, 2, -1, 3, -1},
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "right-skewed",
			tree: []int{1, -1, 2, -1, 3},
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "complete tree",
			tree: []int{1, 2, 3, 4, 5, 6, 7},
			want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name: "asymmetric tree",
			tree: []int{3, 9, 20, -1, -1, 15, 7},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.tree)
			got := LevelOrder(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
