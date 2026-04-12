package heaps

// FindKthLargest returns the kth largest element in the array.
// k is 1-indexed (k=1 means the largest).
func FindKthLargest(nums []int, k int) int {
	// Your solution here

	minHeap := NewMinHeap[int]()

	for _, val := range nums {
		if minHeap.Size() < k {
			minHeap.Insert(Entry[int]{
				Priority: val,
				Value:    val,
			})
			continue
		}

		peeked, _ := minHeap.Peek()

		if val > peeked {
			minHeap.Extract()
			minHeap.Insert(Entry[int]{
				Priority: val,
				Value:    val,
			})
		}
	}

	extracted, _ := minHeap.Extract()

	return extracted
}
