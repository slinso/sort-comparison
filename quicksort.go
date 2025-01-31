package sortcomparison

import "math/rand"

// QuickSort implements the quicksort algorithm which:
// 1. Selects a pivot element from the array
// 2. Partitions other elements into two sub-arrays according to whether they are
//    less than or greater than the pivot
// 3. Recursively sorts the sub-arrays
//
// Time Complexity:
// - Average Case: O(n log n)
// - Best Case: O(n log n)
// - Worst Case: O(n²) - can be mitigated with random pivot selection
//
// Space Complexity: O(log n) average, O(n) worst case
// In-place: Yes
// Stable: No
// Parallel: Yes (with modifications)

func QuickSort[T any](arr []T, less func(a, b T) bool) {
	if len(arr) <= 1 {
		return
	}
	quicksort(arr, 0, len(arr)-1, less)
}

func quicksort[T any](arr []T, low, high int, less func(a, b T) bool) {
	if low < high {
		// Choose random pivot to avoid worst case for sorted arrays
		pivotIndex := low + rand.Intn(high-low+1)
		arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

		// Partition and get pivot position
		pi := quicksortPartition(arr, low, high, less)

		// Recursively sort elements before and after partition
		quicksort(arr, low, pi-1, less)
		quicksort(arr, pi+1, high, less)
	}
}

func quicksortPartition[T any](arr []T, low, high int, less func(a, b T) bool) int {
	pivot := arr[high]
	i := low - 1 // Index of smaller element

	for j := low; j < high; j++ {
		// If current element is smaller than pivot
		if less(arr[j], pivot) {
			i++ // Increment index of smaller element
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Place pivot in its correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
