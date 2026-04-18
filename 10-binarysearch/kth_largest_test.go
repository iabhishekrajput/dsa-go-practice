package binarysearch

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"basic case", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"with duplicates", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"k equals length", []int{3, 2, 1, 5, 6, 4}, 6, 1},
		{"k is 1 (max)", []int{3, 2, 1, 5, 6, 4}, 1, 6},
		{"single element", []int{1}, 1, 1},
		{"two elements k=1", []int{2, 1}, 1, 2},
		{"two elements k=2", []int{2, 1}, 2, 1},
		{"all same", []int{5, 5, 5, 5}, 2, 5},
		{"negative numbers", []int{-1, -2, -3, -4}, 2, -2},
		{"mixed signs", []int{-1, 2, 0, -3, 5}, 3, 0},
		{"sorted ascending", []int{1, 2, 3, 4, 5}, 3, 3},
		{"sorted descending", []int{5, 4, 3, 2, 1}, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// copy to avoid test cross-contamination if solution mutates
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			got := FindKthLargest(nums, tt.k)
			if got != tt.want {
				t.Errorf("FindKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
