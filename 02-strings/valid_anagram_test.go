package strings

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
		{"listen", "silent", true},
		{"hello", "world", false},
		{"aabb", "abab", true},
		{"abc", "abcd", false},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s_%s", tt.s, tt.t)
		t.Run(name, func(t2 *testing.T) {
			got := IsAnagram(tt.s, tt.t)
			if got != tt.want {
				t2.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
