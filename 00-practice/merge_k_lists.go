package practice

// MergeKLists merges k sorted linked lists into one sorted linked list.
func MergeKLists(lists []*ListNode) *ListNode {
	// Your solution here
	if len(lists) == 0 {
		return nil
	}

	heap := NewMinHeap[*ListNode]()

	for _, list := range lists {
		if list != nil {
			heap.Insert(Entry[*ListNode]{Priority: list.Val, Value: list})
		}
	}

	dummyNode := &ListNode{}
	currentNode := dummyNode

	for heap.Size() > 0 {
		extracted, ok := heap.Extract()
		if ok {
			currentNode.Next = extracted
			currentNode = currentNode.Next
			if extracted.Next != nil {
				heap.Insert(Entry[*ListNode]{Priority: extracted.Next.Val, Value: extracted.Next})
			}
		}
	}

	return dummyNode.Next
}
