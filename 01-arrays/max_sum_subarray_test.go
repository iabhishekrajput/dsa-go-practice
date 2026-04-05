package arrays

import (
	"fmt"
	"testing"
)

func TestMaxSumSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 1, 5, 1, 3, 2}, 3, 9},
		{[]int{1, 2, 3, 4, 5}, 2, 9},
		{[]int{5}, 1, 5},
		{[]int{3, 3, 3, 3}, 2, 6},
		{[]int{-1, -2, -3, -4}, 2, -3},
		{[]int{1, -1, 5, -2, 3}, 4, 5},
		{[]int{10, 20, 30, 40, 50}, 5, 150},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("nums=%v,k=%d", tt.nums, tt.k)
		t.Run(name, func(t *testing.T) {
			got := MaxSumSubarray(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("MaxSumSubarray(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
