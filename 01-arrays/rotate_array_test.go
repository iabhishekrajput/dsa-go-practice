package arrays

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "basic rotation",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "rotate by 1",
			nums: []int{1, 2, 3},
			k:    1,
			want: []int{3, 1, 2},
		},
		{
			name: "k equals length (no change)",
			nums: []int{1, 2, 3},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "k greater than length",
			nums: []int{1, 2, 3, 4},
			k:    6,
			want: []int{3, 4, 1, 2},
		},
		{
			name: "rotate by 0",
			nums: []int{1, 2, 3},
			k:    0,
			want: []int{1, 2, 3},
		},
		{
			name: "two elements",
			nums: []int{-1, -100},
			k:    1,
			want: []int{-100, -1},
		},
		{
			name: "single element",
			nums: []int{42},
			k:    5,
			want: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			RotateArray(nums, tt.k)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("RotateArray(%v, %d) = %v, want %v", tt.nums, tt.k, nums, tt.want)
			}
		})
	}
}
