package stacks

import (
	"fmt"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{60, 50, 40, 30}, []int{0, 0, 0, 0}},
		{[]int{50}, []int{0}},
		{[]int{80, 80, 80}, []int{0, 0, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.input)
		t.Run(name, func(t *testing.T) {
			got := DailyTemperatures(tt.input)
			if len(got) != len(tt.want) {
				t.Fatalf("DailyTemperatures(%v) = %v, want %v", tt.input, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("DailyTemperatures(%v) = %v, want %v", tt.input, got, tt.want)
					break
				}
			}
		})
	}
}
