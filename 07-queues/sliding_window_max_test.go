package queues

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{9, 11}, 2, []int{11}},
		{[]int{4, 3, 2, 1}, 2, []int{4, 3, 2}},
		{[]int{1, 2, 3, 4}, 2, []int{2, 3, 4}},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("nums=%v,k=%d", tt.nums, tt.k)
		t.Run(name, func(t *testing.T) {
			got := MaxSlidingWindow(tt.nums, tt.k)
			if len(got) != len(tt.want) {
				t.Fatalf("MaxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("MaxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
					break
				}
			}
		})
	}
}
