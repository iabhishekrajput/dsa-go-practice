package arrays

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		wantK    int
		wantNums []int
	}{
		{[]int{1, 1, 2, 3, 3, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{1, 1, 1}, 1, []int{1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{}, 0, []int{}},
		{[]int{42}, 1, []int{42}},
		{[]int{-3, -1, -1, 0, 0, 0, 5, 5}, 4, []int{-3, -1, 0, 5}},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.input)
		// make a copy so we can show original in failure messages
		input := make([]int, len(tt.input))
		copy(input, tt.input)

		t.Run(name, func(t *testing.T) {
			k := RemoveDuplicates(input)
			if k != tt.wantK {
				t.Errorf("RemoveDuplicates(%v) = %d, want %d", tt.input, k, tt.wantK)
				return
			}
			for i := 0; i < k; i++ {
				if input[i] != tt.wantNums[i] {
					t.Errorf("After RemoveDuplicates(%v): nums[%d] = %d, want %d",
						tt.input, i, input[i], tt.wantNums[i])
				}
			}
		})
	}
}
