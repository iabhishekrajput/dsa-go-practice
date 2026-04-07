package linkedlists

// HasCycle returns true if the linked list contains a cycle.
//
// Constraints:
//   - O(n) time, O(1) space
//
// Hint: Floyd's Tortoise and Hare.
//       slow moves 1 step, fast moves 2 steps.
//       If they meet → cycle. If fast hits nil → no cycle.
func HasCycle(head *ListNode) bool {
	// YOUR CODE HERE

	if head == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
