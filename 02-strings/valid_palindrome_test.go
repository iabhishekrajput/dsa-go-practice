package strings

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{" ", true},
		{"a", true},
		{"ab", false},
		{"aba", true},
		{"0P", false},
		{"Was it a car or a cat I saw?", true},
		{"No 'x' in Nixon", true},
		{"12321", true},
		{"123a321", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
