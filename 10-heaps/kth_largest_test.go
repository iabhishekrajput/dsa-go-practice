package heaps

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "basic example",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "with duplicates",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "k equals 1 (max)",
			nums: []int{7, 3, 5, 1},
			k:    1,
			want: 7,
		},
		{
			name: "k equals length (min)",
			nums: []int{7, 3, 5, 1},
			k:    4,
			want: 1,
		},
		{
			name: "single element",
			nums: []int{42},
			k:    1,
			want: 42,
		},
		{
			name: "all same values",
			nums: []int{5, 5, 5, 5},
			k:    2,
			want: 5,
		},
		{
			name: "negative numbers",
			nums: []int{-1, -5, -3, -2, -4},
			k:    2,
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindKthLargest(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("FindKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
