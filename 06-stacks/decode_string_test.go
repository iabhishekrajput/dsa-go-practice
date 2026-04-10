package stacks

import "testing"

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple repeat",
			input: "3[a]",
			want:  "aaa",
		},
		{
			name:  "nested",
			input: "3[a2[c]]",
			want:  "accaccacc",
		},
		{
			name:  "adjacent groups",
			input: "2[a]3[b]",
			want:  "aabbb",
		},
		{
			name:  "no encoding",
			input: "abc",
			want:  "abc",
		},
		{
			name:  "multi-digit number",
			input: "12[a]",
			want:  "aaaaaaaaaaaa",
		},
		{
			name:  "nested three levels",
			input: "2[a2[b3[c]]]",
			want:  "abcccbcccabcccbccc",
		},
		{
			name:  "mixed plain and encoded",
			input: "x2[ab]y",
			want:  "xababy",
		},
		{
			name:  "single char repeated",
			input: "1[z]",
			want:  "z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DecodeString(tt.input)
			if got != tt.want {
				t.Errorf("DecodeString(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
