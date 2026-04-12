package trees

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name     string
		input    *TreeNode
		expected *TreeNode
	}{
		{
			name: "standard tree",
			input: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			name: "two nodes",
			input: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 1},
			},
			expected: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 1},
			},
		},
		{
			name:     "nil tree",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single node",
			input:    &TreeNode{Val: 1},
			expected: &TreeNode{Val: 1},
		},
		{
			name: "left-skewed becomes right-skewed",
			input: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 1},
				},
			},
			expected: &TreeNode{Val: 3,
				Right: &TreeNode{Val: 2,
					Right: &TreeNode{Val: 1},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InvertTree(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("InvertTree() = %v, want %v", serialize(got), serialize(tt.expected))
			}
		})
	}
}

// serialize does a level-order traversal for readable error output
func serialize(root *TreeNode) []any {
	if root == nil {
		return nil
	}
	result := []any{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, nil)
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	// Trim trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	return result
}
