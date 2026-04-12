package linkedlists

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{5, 10, 15}, []int{1, 8, 20}, []int{1, 5, 8, 10, 15, 20}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v+%v", tt.list1, tt.list2)
		t.Run(name, func(t *testing.T) {
			l1 := buildList(tt.list1)
			l2 := buildList(tt.list2)
			got := listToSlice(MergeTwoLists(l1, l2))

			if len(got) != len(tt.want) {
				t.Fatalf("MergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("MergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, got, tt.want)
					break
				}
			}
		})
	}
}
