package binarysearch

// BinarySearch returns the index of target in a sorted array, or -1 if not found.
func BinarySearch(nums []int, target int) int {
	// Your solution here

	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}
