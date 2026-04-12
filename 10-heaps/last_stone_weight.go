package heaps

// LastStoneWeight simulates smashing the two heaviest stones repeatedly.
// Returns the weight of the last remaining stone, or 0 if none remain.
func LastStoneWeight(stones []int) int {
	// Your solution here

	heap := NewMaxHeap[int]()
	for _, stone := range stones {
		heap.Insert(Entry[int]{Priority: stone, Value: stone})
	}

	for heap.Size() > 1 {
		stone1, _ := heap.Extract()
		stone2, _ := heap.Extract()

		remianing := stone1 - stone2
		if remianing > 0 {
			heap.Insert(Entry[int]{Priority: remianing, Value: remianing})
		}
	}

	extracted, ok := heap.Peek()
	if ok {
		return extracted
	}

	return 0
}
