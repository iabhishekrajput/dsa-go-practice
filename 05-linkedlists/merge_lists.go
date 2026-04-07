package linkedlists

// MergeTwoLists merges two sorted linked lists into one sorted list.
//
// Example:
//   1→2→4 + 1→3→4 = 1→1→2→3→4→4
//
// Constraints:
//   - O(n + m) time, O(1) space (reuse existing nodes)
//
// Hint: Use a dummy node as the starting point.
//       Compare heads of both lists, attach the smaller one.
//       When one runs out, attach the remainder of the other.
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// YOUR CODE HERE

	dummy := &ListNode{}
	node1, node2 := list1, list2
	curr := dummy

	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			curr.Next = node1
			node1 = node1.Next
		} else {
			curr.Next = node2
			node2 = node2.Next
		}

		curr = curr.Next
	}

	if node1 != nil {
		curr.Next = node1
	} else if node2 != nil {
		curr.Next = node2
	}

	return dummy.Next
}
