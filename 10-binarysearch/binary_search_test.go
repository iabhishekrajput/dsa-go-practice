package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "found in middle",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "not found",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			name:     "first element",
			nums:     []int{1, 3, 5, 7},
			target:   1,
			expected: 0,
		},
		{
			name:     "last element",
			nums:     []int{1, 3, 5, 7},
			target:   7,
			expected: 3,
		},
		{
			name:     "single element found",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "single element not found",
			nums:     []int{5},
			target:   3,
			expected: -1,
		},
		{
			name:     "empty array",
			nums:     []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "negative numbers",
			nums:     []int{-10, -5, -3, 0, 2},
			target:   -3,
			expected: 2,
		},
		{
			name:     "target smaller than all",
			nums:     []int{2, 4, 6, 8},
			target:   1,
			expected: -1,
		},
		{
			name:     "target larger than all",
			nums:     []int{2, 4, 6, 8},
			target:   10,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
