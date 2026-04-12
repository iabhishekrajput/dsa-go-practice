package trees

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "standard tree",
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3,
					Right: &TreeNode{Val: 4},
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			name: "left side deeper",
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name:     "nil tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "single node",
			root:     &TreeNode{Val: 42},
			expected: []int{42},
		},
		{
			name: "complete tree",
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: []int{1, 3, 7},
		},
		{
			name: "zigzag deeper on left",
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 4,
						Right: &TreeNode{Val: 7},
					},
				},
				Right: &TreeNode{Val: 3,
					Right: &TreeNode{Val: 5},
				},
			},
			expected: []int{1, 3, 5, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RightSideView(tt.root)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("RightSideView() = %v, want %v", got, tt.expected)
			}
		})
	}
}
