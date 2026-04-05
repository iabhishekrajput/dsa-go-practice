package arrays

import (
	"fmt"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{[]int{1, 5, 8, 3, 9, 2}, 11, []int{2, 3}},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("nums=%v,target=%d", tt.nums, tt.target)
		t.Run(name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if got == nil || len(got) != 2 {
				t.Fatalf("TwoSum(%v, %d) = %v, want 2 indices", tt.nums, tt.target, got)
			}
			// sort both for comparison (order of indices doesn't matter)
			sort.Ints(got)
			sort.Ints(tt.want)
			if got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
