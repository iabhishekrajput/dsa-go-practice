package strings

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"single", "single"},
		{"  spaces  everywhere  ", "everywhere spaces"},
		{"Alice does not even like bob", "bob like even not does Alice"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ReverseWords(tt.input)
			if got != tt.want {
				t.Errorf("ReverseWords(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
