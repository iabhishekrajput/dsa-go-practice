package practice

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{4, 0, 5, 0, 0, 6}, []int{4, 5, 6, 0, 0, 0}},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.input)
		input := make([]int, len(tt.input))
		copy(input, tt.input)

		t.Run(name, func(t *testing.T) {
			MoveZeroes(input)
			for i := range input {
				if input[i] != tt.want[i] {
					t.Errorf("MoveZeroes(%v) = %v, want %v", tt.input, input, tt.want)
					break
				}
			}
		})
	}
}
