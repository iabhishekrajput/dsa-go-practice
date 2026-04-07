package hashmaps

import (
	"fmt"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{4, 4, 4, 1, 1, 2, 2, 2, 3}, 2, []int{2, 4}},
		{[]int{5, 5, 5, 5}, 1, []int{5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("nums=%v,k=%d", tt.nums, tt.k)
		t.Run(name, func(t *testing.T) {
			got := TopKFrequent(tt.nums, tt.k)

			if len(got) != tt.k {
				t.Fatalf("TopKFrequent(%v, %d) returned %d elements, want %d", tt.nums, tt.k, len(got), tt.k)
			}

			sort.Ints(got)
			sort.Ints(tt.want)

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("TopKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
					break
				}
			}
		})
	}
}
