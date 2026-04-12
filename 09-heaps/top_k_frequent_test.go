package heaps

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "basic example",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "single element",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "all same frequency",
			nums: []int{1, 2, 3},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "k equals 1",
			nums: []int{4, 4, 4, 1, 2, 2},
			k:    1,
			want: []int{4},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, 2, 2, 2, 3},
			k:    2,
			want: []int{-1, 2},
		},
		{
			name: "larger example",
			nums: []int{5, 3, 1, 1, 1, 3, 5, 73, 1},
			k:    2,
			want: []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.nums, tt.k)
			sort.Ints(got)
			sort.Ints(tt.want)
			if len(got) != len(tt.want) {
				t.Errorf("TopKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("TopKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
					return
				}
			}
		})
	}
}
