package linkedlists

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.input)
		t.Run(name, func(t *testing.T) {
			head := buildList(tt.input)
			got := listToSlice(ReverseList(head))

			if len(got) != len(tt.want) {
				t.Fatalf("ReverseList(%v) = %v, want %v", tt.input, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ReverseList(%v) = %v, want %v", tt.input, got, tt.want)
					break
				}
			}
		})
	}
}
