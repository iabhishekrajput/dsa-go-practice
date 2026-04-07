package hashmaps

import "sort"

// TopKFrequent returns the k most frequent elements from nums.
// The output order does not matter. It is guaranteed that the answer is unique.
//
// Example:
//   TopKFrequent([]int{1,1,1,2,2,3}, 2) → [1, 2]
//
// Constraints:
//   - O(n log n) time is fine for now (we'll learn O(n) with heaps later)
//
// Hint: Step 1 — build a frequency map: map[int]int
//       Step 2 — collect the keys, sort them by frequency (descending)
//       Step 3 — return the first k keys
func TopKFrequent(nums []int, k int) []int {
	// YOUR CODE HERE

	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}

	keys := make([]int, len(freq))
	for num := range freq {
		keys = append(keys, num)
	}

	sort.Slice(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	return keys[:k]
}
