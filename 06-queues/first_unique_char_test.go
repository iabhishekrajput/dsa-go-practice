package queues

import "testing"

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
		{"a", 0},
		{"aadadaad", -1},
		{"dddccdbba", 8},
		{"abcabc", -1},
		{"abcdef", 0},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := FirstUniqChar(tt.input)
			if got != tt.want {
				t.Errorf("FirstUniqChar(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
