package linkedlists

// ReverseList reverses a singly linked list and returns the new head.
//
// Example:
//   1 → 2 → 3 → 4 → 5  becomes  5 → 4 → 3 → 2 → 1
//
// Constraints:
//   - O(n) time, O(1) space
//
// Hint: Three pointers — prev, curr, next.
//       At each step: save next, flip curr.Next to prev, advance both.
func ReverseList(head *ListNode) *ListNode {
	// YOUR CODE HERE

	if head == nil {
		return head
	}

	var prev *ListNode = nil
	curr, next := head, head.Next

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
