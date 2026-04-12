package heaps

// TopKFrequent returns the k most frequent elements in nums.
// The result can be in any order.
func TopKFrequent(nums []int, k int) []int {
	// Your solution here

	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}

	heap := NewMinHeap[int]()
	for e, f := range freq {
		if heap.Size() < k {
			heap.Insert(Entry[int]{Priority: f, Value: e})
			continue
		}

		peeked, _ := heap.PeekPriority()
		if f > peeked {
			heap.Extract()
			heap.Insert(Entry[int]{Priority: f, Value: e})
		}
	}

	result := []int{}
	for heap.Size() > 0 {
		extracted, _ := heap.Extract()
		result = append(result, extracted)
	}

	return result
}
