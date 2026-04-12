package stacks

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		name      string
		asteroids []int
		want      []int
	}{
		{
			name:      "small one destroyed",
			asteroids: []int{5, 10, -5},
			want:      []int{5, 10},
		},
		{
			name:      "equal size both destroyed",
			asteroids: []int{8, -8},
			want:      []int{},
		},
		{
			name:      "bigger negative survives",
			asteroids: []int{10, 2, -5},
			want:      []int{10},
		},
		{
			name:      "chain reaction",
			asteroids: []int{5, 3, -7},
			want:      []int{-7},
		},
		{
			name:      "no collision same direction",
			asteroids: []int{1, 2, 3},
			want:      []int{1, 2, 3},
		},
		{
			name:      "no collision all left",
			asteroids: []int{-1, -2, -3},
			want:      []int{-1, -2, -3},
		},
		{
			name:      "no collision left then right",
			asteroids: []int{-2, -1, 1, 2},
			want:      []int{-2, -1, 1, 2},
		},
		{
			name:      "multiple collisions",
			asteroids: []int{1, 2, 3, -3, -2, -1},
			want:      []int{},
		},
		{
			name:      "negative survives through all",
			asteroids: []int{1, 2, 3, -10},
			want:      []int{-10},
		},
		{
			name:      "single asteroid",
			asteroids: []int{42},
			want:      []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AsteroidCollision(tt.asteroids)
			if len(got) == 0 && len(tt.want) == 0 {
				return // both empty — pass
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsteroidCollision(%v) = %v, want %v", tt.asteroids, got, tt.want)
			}
		})
	}
}
