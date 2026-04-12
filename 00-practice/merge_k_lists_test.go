package practice

import (
	"reflect"
	"testing"
)

func toList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func toSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		{
			name:     "three sorted lists",
			lists:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "empty input",
			lists:    [][]int{},
			expected: nil,
		},
		{
			name:     "single empty list",
			lists:    [][]int{{}},
			expected: nil,
		},
		{
			name:     "one list",
			lists:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "lists with single elements",
			lists:    [][]int{{5}, {3}, {1}, {4}, {2}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "some nil lists mixed in",
			lists:    [][]int{{}, {1, 3}, {}, {2, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "all same values",
			lists:    [][]int{{2, 2}, {2, 2}, {2}},
			expected: []int{2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lists []*ListNode
			for _, vals := range tt.lists {
				lists = append(lists, toList(vals))
			}
			got := MergeKLists(lists)
			result := toSlice(got)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeKLists() = %v, want %v", result, tt.expected)
			}
		})
	}
}
