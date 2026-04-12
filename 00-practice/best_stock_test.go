package practice

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 0},
		{[]int{3, 3, 3}, 0},
		{[]int{1}, 0},
		{[]int{2, 4, 1, 7}, 6},
		{[]int{1, 2, 3, 4, 5}, 4},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.prices)
		t.Run(name, func(t *testing.T) {
			got := MaxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
