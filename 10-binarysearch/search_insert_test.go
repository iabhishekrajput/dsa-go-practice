package binarysearch

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "found exact",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "insert in middle",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "insert at end",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "insert at start",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "single element found",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "single element insert before",
			nums:     []int{3},
			target:   1,
			expected: 0,
		},
		{
			name:     "single element insert after",
			nums:     []int{3},
			target:   5,
			expected: 1,
		},
		{
			name:     "empty array",
			nums:     []int{},
			target:   5,
			expected: 0,
		},
		{
			name:     "insert between duplicates area",
			nums:     []int{1, 3, 5, 5, 6},
			target:   4,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInsert(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("SearchInsert(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
