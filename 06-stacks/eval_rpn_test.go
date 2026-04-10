package stacks

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "simple addition",
			tokens: []string{"2", "1", "+"},
			want:   3,
		},
		{
			name:   "addition then multiply",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "subtraction order matters",
			tokens: []string{"10", "3", "-"},
			want:   7,
		},
		{
			name:   "division truncates toward zero",
			tokens: []string{"7", "2", "/"},
			want:   3,
		},
		{
			name:   "negative division truncates toward zero",
			tokens: []string{"7", "-2", "/"},
			want:   -3,
		},
		{
			name:   "complex expression",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			name:   "leetcode example 3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
		{
			name:   "single number",
			tokens: []string{"42"},
			want:   42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvalRPN(tt.tokens)
			if got != tt.want {
				t.Errorf("EvalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
			}
		})
	}
}
