package sortcomparison

import "math/bits"

func Introsort[T any](arr []T, less func(a, b T) bool) {
	maxDepth := 2 * bits.Len(uint(len(arr)))
	introSortRec(arr, maxDepth, less)
}

func introSortRec[T any](arr []T, maxDepth int, less func(a, b T) bool) {
	n := len(arr)

	// Use insertion sort for small arrays
	if n <= 16 {
		insertionSort(arr, less)
		return
	}

	// Switch to heapsort if depth limit exceeded
	if maxDepth == 0 {
		Heapsort(arr, less)
		return
	}

	// Otherwise, do quicksort
	pivot := partition(arr, less)
	introSortRec(arr[:pivot], maxDepth-1, less)
	introSortRec(arr[pivot+1:], maxDepth-1, less)
}

func partition[T any](arr []T, less func(a, b T) bool) int {
	pivot := arr[len(arr)-1]
	i := -1

	for j := 0; j < len(arr)-1; j++ {
		if less(arr[j], pivot) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
	return i + 1
}

func insertionSort[T any](arr []T, less func(a, b T) bool) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && less(key, arr[j]) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
