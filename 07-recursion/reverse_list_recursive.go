package recursion

// ListNode is a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseListRecursive reverses a singly linked list using recursion.
// Returns the new head of the reversed list.
func ReverseListRecursive(head *ListNode) *ListNode {
	// Your solution here
	
	if head == nil || head.Next == nil {
		return head
	}

	front := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return front
}
