package linkedlists

import "testing"

func TestHasCycle(t *testing.T) {
	// Case 1: cycle at node index 1
	// 3 → 2 → 0 → -4 → (back to 2)
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2 // cycle

	if !HasCycle(n1) {
		t.Error("Expected cycle, got false")
	}

	// Case 2: cycle at node index 0
	// 1 → 2 → (back to 1)
	a1 := &ListNode{Val: 1}
	a2 := &ListNode{Val: 2}
	a1.Next = a2
	a2.Next = a1 // cycle

	if !HasCycle(a1) {
		t.Error("Expected cycle, got false")
	}

	// Case 3: no cycle
	// 1 → 2 → 3 → nil
	if HasCycle(buildList([]int{1, 2, 3})) {
		t.Error("Expected no cycle, got true")
	}

	// Case 4: single node, no cycle
	if HasCycle(buildList([]int{1})) {
		t.Error("Expected no cycle for single node, got true")
	}

	// Case 5: nil
	if HasCycle(nil) {
		t.Error("Expected no cycle for nil, got true")
	}

	// Case 6: single node pointing to itself
	solo := &ListNode{Val: 1}
	solo.Next = solo
	if !HasCycle(solo) {
		t.Error("Expected cycle for self-loop, got false")
	}
}
