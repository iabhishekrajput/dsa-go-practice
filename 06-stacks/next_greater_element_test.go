package stacks

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "basic example",
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2, 5},
			want:  []int{5, 3, 5},
		},
		{
			name:  "no greater elements",
			nums1: []int{4, 3},
			nums2: []int{1, 2, 3, 4},
			want:  []int{-1, 4},
		},
		{
			name:  "all have next greater",
			nums1: []int{1, 2, 3},
			nums2: []int{1, 2, 3, 4},
			want:  []int{2, 3, 4},
		},
		{
			name:  "single element",
			nums1: []int{5},
			nums2: []int{5},
			want:  []int{-1},
		},
		{
			name:  "nums1 equals nums2",
			nums1: []int{3, 1, 2},
			nums2: []int{3, 1, 2},
			want:  []int{-1, 2, -1},
		},
		{
			name:  "decreasing then spike",
			nums1: []int{5, 4, 3},
			nums2: []int{5, 4, 3, 10},
			want:  []int{10, 10, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextGreaterElement(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterElement(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
