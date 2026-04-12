package sorting

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "mixed",
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "three elements",
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			name:     "already sorted",
			input:    []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "reverse sorted",
			input:    []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "all twos",
			input:    []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "two elements",
			input:    []int{1, 0},
			expected: []int{0, 1},
		},
		{
			name:     "only zeros and twos",
			input:    []int{2, 0, 2, 0, 0, 2},
			expected: []int{0, 0, 0, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.input))
			copy(nums, tt.input)
			SortColors(nums)
			if !reflect.DeepEqual(nums, tt.expected) {
				t.Errorf("SortColors(%v) = %v, want %v", tt.input, nums, tt.expected)
			}
		})
	}
}
