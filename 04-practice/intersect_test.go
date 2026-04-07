package practice

import (
	"fmt"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{[]int{1}, []int{1}, []int{1}},
		{[]int{3, 1, 2}, []int{1, 1}, []int{1}},
		{[]int{1, 1, 1, 2}, []int{1, 1, 3}, []int{1, 1}},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v_%v", tt.nums1, tt.nums2)
		t.Run(name, func(t *testing.T) {
			got := Intersect(tt.nums1, tt.nums2)
			sort.Ints(got)
			sort.Ints(tt.want)

			if len(got) != len(tt.want) {
				t.Fatalf("Intersect(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Intersect(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
					break
				}
			}
		})
	}
}
