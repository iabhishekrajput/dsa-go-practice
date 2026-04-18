package binarysearch

import "math/rand"

// FindKthLargest returns the kth largest element in nums using quickselect.
// Expected O(n) average time, O(n^2) worst case.
// You may modify nums.
func FindKthLargest(nums []int, k int) int {
	// Your solution here
	quickSortHelper(nums, 0, len(nums)-1, k)
	return nums[(len(nums) - k)]
}

func quickSortHelper(nums []int, low, high, k int) {
	if low < high {
		pIndex := partition(nums, low, high)

		if pIndex == len(nums)-k {
			return
		}

		if pIndex < len(nums)-k {
			quickSortHelper(nums, pIndex+1, high, k)
		} else {
			quickSortHelper(nums, low, pIndex-1, k)
		}
	}
}

func partition(nums []int, low, high int) int {
	pivotIndex := rand.Int()%(high-low+1) + low
	nums[low], nums[pivotIndex] = nums[pivotIndex], nums[low]
	pivot, i, j := nums[low], low, high

	for i < j {
		for nums[i] <= pivot && i < high {
			i++
		}

		for pivot < nums[j] && low < j {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[low], nums[j] = nums[j], nums[low]

	return j
}
