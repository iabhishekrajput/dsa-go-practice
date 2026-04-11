package recursion

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "basic unsorted",
			input: []int{38, 27, 43, 3, 9, 82, 10},
			want:  []int{3, 9, 10, 27, 38, 43, 82},
		},
		{
			name:  "already sorted",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "reverse sorted",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "duplicates",
			input: []int{4, 2, 4, 1, 2},
			want:  []int{1, 2, 2, 4, 4},
		},
		{
			name:  "single element",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "two elements",
			input: []int{5, 3},
			want:  []int{3, 5},
		},
		{
			name:  "empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "negative numbers",
			input: []int{-3, 5, -1, 0, 2, -8},
			want:  []int{-8, -3, -1, 0, 2, 5},
		},
		{
			name:  "all same",
			input: []int{7, 7, 7, 7},
			want:  []int{7, 7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.input)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
