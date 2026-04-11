package recursion

import (
	"reflect"
	"testing"
)

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestReverseListRecursive(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "multiple elements",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "two elements",
			input: []int{1, 2},
			want:  []int{2, 1},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "empty list",
			input: []int{},
			want:  nil,
		},
		{
			name:  "three elements",
			input: []int{10, 20, 30},
			want:  []int{30, 20, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			got := listToSlice(ReverseListRecursive(head))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseListRecursive(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
