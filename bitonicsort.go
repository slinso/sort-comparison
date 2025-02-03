package sortcomparison

import "math"

// BitonicSort implements an optimized bitonic sorting network.
// Best for parallel hardware; performs O(n log²n) comparisons.
// only working for power of 2 arrays
func BitonicSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Find next power of 2 length
	n := len(arr)

	bitonicSort(arr, 0, n, true)
}

// BitonicSort sorts the entire array in ascending order
func BitonicSortAny(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}

	// Pad the array to the next power of two
	nextPow2 := nextPowerOfTwo(n)
	paddedArr := make([]int, nextPow2)
	copy(paddedArr, arr)
	for i := n; i < nextPow2; i++ {
		paddedArr[i] = math.MaxInt32 // Use a large value as padding
	}

	// Sort the padded array
	bitonicSort(paddedArr, 0, nextPow2, true)

	// Copy the sorted elements back to the original array
	copy(arr, paddedArr[:n])
}

// bitonicMerge recursively sorts a bitonic sequence in ascending or descending order
func bitonicMerge(arr []int, low, cnt int, ascending bool) {
	if cnt > 1 {
		k := cnt / 2
		for i := low; i < low+k; i++ {
			compareAndSwap(arr, i, i+k, ascending)
		}
		bitonicMerge(arr, low, k, ascending)
		bitonicMerge(arr, low+k, k, ascending)
	}
}

// bitonicSort recursively sorts the array in ascending or descending order
func bitonicSort(arr []int, low, cnt int, ascending bool) {
	if cnt > 1 {
		k := cnt / 2
		bitonicSort(arr, low, k, true)
		bitonicSort(arr, low+k, k, false)
		bitonicMerge(arr, low, cnt, ascending)
	}
}
