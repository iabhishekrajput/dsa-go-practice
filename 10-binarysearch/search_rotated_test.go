package binarysearch

import "testing"

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"found in right half", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"not found", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"single element found", []int{1}, 1, 0},
		{"single element not found", []int{1}, 0, -1},
		{"found in left half", []int{4, 5, 6, 7, 0, 1, 2}, 5, 1},
		{"target is pivot", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{"target is last", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{"not rotated", []int{1, 2, 3, 4, 5}, 3, 2},
		{"rotated by one", []int{2, 3, 4, 5, 1}, 1, 4},
		{"two elements rotated", []int{2, 1}, 1, 1},
		{"two elements not found", []int{2, 1}, 3, -1},
		{"large rotation", []int{6, 7, 0, 1, 2, 3, 4, 5}, 7, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRotated(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("SearchRotated(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
