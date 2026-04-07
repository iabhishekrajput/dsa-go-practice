package practice

// Intersect returns the intersection of two slices.
// Each element appears as many times as it shows in both arrays.
//
// Example:
//   Intersect([]int{1,2,2,1}, []int{2,2}) → [2,2]
//   Intersect([]int{4,9,5}, []int{9,4,9,8,4}) → [4,9]
//
// Constraints:
//   - O(n + m) time
func Intersect(nums1 []int, nums2 []int) []int {
	// YOUR CODE HERE

	freqMap := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		freqMap[num]++
	}

	var result []int
	for _, num := range nums2 {
		if freqMap[num] > 0 {
			result = append(result, num)
			freqMap[num]--
		}
	}

	return result
}
