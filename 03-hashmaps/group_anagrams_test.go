package hashmaps

import (
	"fmt"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input []string
		want  [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}},
		},
		{
			[]string{""},
			[][]string{{""}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
		{
			[]string{"abc", "bca", "cab", "xyz", "zyx"},
			[][]string{{"abc", "bca", "cab"}, {"xyz", "zyx"}},
		},
		{
			[]string{"no", "on", "is", "si"},
			[][]string{{"is", "si"}, {"no", "on"}},
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.input)
		t.Run(name, func(t *testing.T) {
			got := GroupAnagrams(tt.input)

			// Sort everything for comparison
			for _, g := range got {
				sort.Strings(g)
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i][0] < got[j][0]
			})

			if len(got) != len(tt.want) {
				t.Fatalf("GroupAnagrams(%v) returned %d groups, want %d\ngot:  %v", tt.input, len(got), len(tt.want), got)
			}

			for i := range got {
				if len(got[i]) != len(tt.want[i]) {
					t.Errorf("Group %d: got %v, want %v", i, got[i], tt.want[i])
					continue
				}
				for j := range got[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("Group %d, item %d: got %q, want %q", i, j, got[i][j], tt.want[i][j])
					}
				}
			}
		})
	}
}
