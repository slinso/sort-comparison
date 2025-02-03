package sortcomparison

import (
	"golang.org/x/exp/constraints"
)

// insertionSort performs an in-place insertion sort on the slice.
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// insertionSortLowHigh is an in-place insertion sort that operates on arr[low:high+1].
func insertionSortLowHigh(arr []int, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1
		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// compareAndSwap swaps the elements if they are not in the desired order
func compareAndSwap(arr []int, i, j int, ascending bool) {
	if (arr[i] > arr[j] && ascending) || (arr[i] < arr[j] && !ascending) {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
