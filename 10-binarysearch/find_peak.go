package binarysearch

// FindPeakElement returns the index of any peak element.
// A peak is strictly greater than its neighbors.
// nums[-1] and nums[n] are considered -infinity.
// Must run in O(log n) time.
func FindPeakElement(nums []int) int {
	// Your solution here

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
