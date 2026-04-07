package hashmaps

// LongestConsecutive returns the length of the longest consecutive
// element sequence in an unsorted slice.
//
// Example:
//   LongestConsecutive([]int{100, 4, 200, 1, 3, 2}) → 4
//
// Constraints:
//   - O(n) time (no sorting!)
//
// Hint: Build a set (map[int]bool). For each number, if num-1 is NOT
//       in the set, it's the start of a sequence. Count forward from there.
func LongestConsecutive(nums []int) int {
	// YOUR CODE HERE

	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}

	maxLen := 0
	for v := range set {
		_, ok := set[v - 1]; if ok {
			set[v] = struct{}{}
			continue
		}

		currLen := 1
		next := v + 1
		for {
			_, ok := set[next]; if !ok {
				break
			}

			currLen++
			next++
		}
		maxLen = max(maxLen, currLen)
	}

	return maxLen
}
