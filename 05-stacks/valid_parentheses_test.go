package stacks

import "testing"

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"(", false},
		{")", false},
		{"((()))", true},
		{"{[()]}", true},
		{"((())", false},
		{"]", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsValidParentheses(tt.input)
			if got != tt.want {
				t.Errorf("IsValidParentheses(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
