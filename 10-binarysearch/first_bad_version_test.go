package binarysearch

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		firstBad int
	}{
		{
			name:     "bad in middle",
			n:        7,
			firstBad: 4,
		},
		{
			name:     "first version is bad",
			n:        5,
			firstBad: 1,
		},
		{
			name:     "last version is bad",
			n:        5,
			firstBad: 5,
		},
		{
			name:     "single version bad",
			n:        1,
			firstBad: 1,
		},
		{
			name:     "two versions first bad",
			n:        2,
			firstBad: 1,
		},
		{
			name:     "two versions second bad",
			n:        2,
			firstBad: 2,
		},
		{
			name:     "large range",
			n:        1000000,
			firstBad: 999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isBadVersion := func(version int) bool {
				return version >= tt.firstBad
			}
			got := FirstBadVersion(tt.n, isBadVersion)
			if got != tt.firstBad {
				t.Errorf("FirstBadVersion(%d) = %d, want %d", tt.n, got, tt.firstBad)
			}
		})
	}
}
