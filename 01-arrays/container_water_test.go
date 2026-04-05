package arrays

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{1, 2, 4, 3}, 4},
		{[]int{2, 3, 10, 5, 7, 8, 9}, 36},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.height)
		t.Run(name, func(t *testing.T) {
			got := MaxArea(tt.height)
			if got != tt.want {
				t.Errorf("MaxArea(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
