package heaps

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{
			name:   "basic example",
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
		{
			name:   "single stone",
			stones: []int{5},
			want:   5,
		},
		{
			name:   "two equal stones",
			stones: []int{3, 3},
			want:   0,
		},
		{
			name:   "two different stones",
			stones: []int{3, 7},
			want:   4,
		},
		{
			name:   "all equal stones even count",
			stones: []int{2, 2, 2, 2},
			want:   0,
		},
		{
			name:   "all equal stones odd count",
			stones: []int{5, 5, 5},
			want:   5,
		},
		{
			name:   "descending weights",
			stones: []int{10, 5, 3, 1},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LastStoneWeight(tt.stones)
			if got != tt.want {
				t.Errorf("LastStoneWeight(%v) = %d, want %d", tt.stones, got, tt.want)
			}
		})
	}
}
