package linkedlists

// ListNode is a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Helper: build a linked list from a slice (for testing)
func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper: convert a linked list to a slice (for testing)
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
