package binarysearch

import "testing"

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		valid []int // any of these indices is acceptable
	}{
		{"single peak in middle", []int{1, 2, 3, 1}, []int{2}},
		{"multiple peaks", []int{1, 2, 1, 3, 5, 6, 4}, []int{1, 5}},
		{"single element", []int{1}, []int{0}},
		{"two elements ascending", []int{1, 2}, []int{1}},
		{"two elements descending", []int{2, 1}, []int{0}},
		{"peak at start", []int{5, 3, 2, 1}, []int{0}},
		{"peak at end", []int{1, 2, 3, 5}, []int{3}},
		{"valley then peak", []int{3, 1, 2}, []int{0, 2}},
		{"long ascending", []int{1, 2, 3, 4, 5}, []int{4}},
		{"long descending", []int{5, 4, 3, 2, 1}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindPeakElement(tt.nums)
			found := false
			for _, v := range tt.valid {
				if got == v {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("FindPeakElement(%v) = %d, want one of %v", tt.nums, got, tt.valid)
			}
		})
	}
}
