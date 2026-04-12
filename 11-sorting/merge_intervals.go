package sorting

import "math/rand"

// MergeIntervals merges all overlapping intervals and returns the result.
// Each interval is [start, end].
func MergeIntervals(intervals [][]int) [][]int {
	// Your solution here
	sortIntervals(intervals, 0, len(intervals)-1)

	result := [][]int{}
	i := 0
	for i < len(intervals) {
		left := intervals[i][0]
		right := intervals[i][1]

		for i < len(intervals)-1 && right >= intervals[i+1][0] {
			right = max(right, intervals[i+1][1])
			i++
		}

		result = append(result, []int{left, right})
		i++
	}

	return result
}

func sortIntervals(arr [][]int, low, high int) {
	if low < high {
		pIndex := parition(arr, low, high)

		sortIntervals(arr, low, pIndex-1)
		sortIntervals(arr, pIndex+1, high)
	}
}

func parition(arr [][]int, low, high int) int {
	pivotIndex := (rand.Int() % (high - low)) + low
	arr[pivotIndex], arr[low] = arr[low], arr[pivotIndex]
	pivot, i, j := arr[low], low, high

	for i < j {
		for arr[i][0] <= pivot[0] && i <= high-1 {
			i++
		}

		for arr[j][0] > pivot[0] && j >= low+1 {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low]

	return j
}
