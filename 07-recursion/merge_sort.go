package recursion

// MergeSort sorts a slice of integers using the merge sort algorithm.
// Returns a new sorted slice.
func MergeSort(nums []int) []int {
	// Your solution here

	if len(nums) < 2 {
		return nums
	}

	left, mid, right := 0, len(nums)/2, len(nums)
	leftSortedArray := MergeSort(nums[left:mid])
	rightSortedArray := MergeSort(nums[mid:right])

	leftPointer, rightPointer := 0, 0
	sortedArray := []int{}
	for leftPointer < len(leftSortedArray) && rightPointer < len(rightSortedArray) {
		if leftSortedArray[leftPointer] < rightSortedArray[rightPointer] {
			sortedArray = append(sortedArray, leftSortedArray[leftPointer])
			leftPointer++
		} else {
			sortedArray = append(sortedArray, rightSortedArray[rightPointer])
			rightPointer++
		}
	}

	if leftPointer < len(leftSortedArray) {
		sortedArray = append(sortedArray, leftSortedArray[leftPointer:]...)
	}
	
	if rightPointer < len(rightSortedArray) {
		sortedArray = append(sortedArray, rightSortedArray[rightPointer:]...)
	}

	return sortedArray
}
