package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "mixed values",
			input:    []int{3, 6, 8, 10, 1, 2, 1},
			expected: []int{1, 1, 2, 3, 6, 8, 10},
		},
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all duplicates",
			input:    []int{3, 3, 3, 3},
			expected: []int{3, 3, 3, 3},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "two elements unsorted",
			input:    []int{5, 1},
			expected: []int{1, 5},
		},
		{
			name:     "negatives",
			input:    []int{-3, 4, -1, 0, 2, -5},
			expected: []int{-5, -3, -1, 0, 2, 4},
		},
		{
			name:     "duplicates mixed",
			input:    []int{4, 2, 4, 1, 2, 1},
			expected: []int{1, 1, 2, 2, 4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.input))
			copy(nums, tt.input)
			QuickSort(nums)
			if !reflect.DeepEqual(nums, tt.expected) {
				t.Errorf("QuickSort(%v) = %v, want %v", tt.input, nums, tt.expected)
			}
		})
	}
}
