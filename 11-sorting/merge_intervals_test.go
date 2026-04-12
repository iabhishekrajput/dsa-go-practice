package sorting

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "overlapping intervals",
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:     "touching boundary",
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "no overlap",
			input:    [][]int{{1, 2}, {4, 5}, {7, 8}},
			expected: [][]int{{1, 2}, {4, 5}, {7, 8}},
		},
		{
			name:     "all merge into one",
			input:    [][]int{{1, 4}, {2, 5}, {3, 6}},
			expected: [][]int{{1, 6}},
		},
		{
			name:     "single interval",
			input:    [][]int{{1, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "unsorted input",
			input:    [][]int{{8, 10}, {1, 3}, {2, 6}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:     "nested interval",
			input:    [][]int{{1, 10}, {3, 5}, {6, 8}},
			expected: [][]int{{1, 10}},
		},
		{
			name:     "identical intervals",
			input:    [][]int{{1, 3}, {1, 3}, {1, 3}},
			expected: [][]int{{1, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeIntervals(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MergeIntervals(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
