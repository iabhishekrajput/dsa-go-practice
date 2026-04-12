package sorting

import "math/rand"

// QuickSort sorts nums in-place using the quick sort algorithm.
func QuickSort(nums []int) {
	// Your solution here
	quickSortHelper(nums, 0, len(nums)-1)
}

func quickSortHelper(nums []int, low, high int) {
	if low < high {
		pIndex := partition(nums, low, high)

		quickSortHelper(nums, low, pIndex-1)
		quickSortHelper(nums, pIndex+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivotIdx := (rand.Int() % (high - low)) + low
	nums[low], nums[pivotIdx] = nums[pivotIdx], nums[low]	
	pivot, i, j := nums[low], low, high
	for i < j {
		for nums[i] <= pivot && i <= high-1 {
			i++
		}

		for nums[j] > pivot && j >= low+1 {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[low], nums[j] = nums[j], nums[low]

	return j
}
