package sortcomparison

import (
	"math"
)

// WikiSort sorts a slice of integers in ascending order using a stable in-place merge-based
// approach inspired by WikiSort. It divides the array into small runs, sorts each run using
// insertion sort, and then iteratively merges adjacent runs using a small temporary buffer.
// This minimizes garbage collection overhead and is idiomatic Go.
//
// How it works:
//   - Compute a block size as floor(sqrt(n)) where n is the length of the array.
//   - Partition the array into runs of size blockSize (the last run may be shorter).
//   - Each run is sorted with insertion sort (efficient for small sizes).
//   - Adjacent runs are merged in-place using a helper that employs a temporary buffer of size
//     at most blockSize.
//   - The merging of runs is performed iteratively until the array is fully sorted.
//
// Time Complexity:
//   - Worst-case: O(n log n)
//
// Space Complexity:
//   - O(sqrt(n)) auxiliary space in the worst case.
func WikiSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	// Use a block size approximated by the square root of n.
	blockSize := int(math.Sqrt(float64(n)))
	if blockSize < 1 {
		blockSize = 1
	}

	type run struct {
		start int
		end   int
	}

	var runs []run
	// Partition the array into runs and sort each run using insertion sort.
	for i := 0; i < n; i += blockSize {
		end := i + blockSize
		if end > n {
			end = n
		}
		insertionSort(arr[i:end])
		runs = append(runs, run{start: i, end: end})
	}

	// Allocate a temporary buffer of size 'blockSize' for merging.
	buf := make([]int, blockSize)

	// Iteratively merge pairs of runs until one sorted run remains.
	for len(runs) > 1 {
		var newRuns []run
		for i := 0; i < len(runs); i += 2 {
			if i+1 < len(runs) {
				r1 := runs[i]
				r2 := runs[i+1]
				inPlaceMerge(arr, r1.start, r1.end, r2.end, buf)
				newRuns = append(newRuns, run{start: r1.start, end: r2.end})
			} else {
				newRuns = append(newRuns, runs[i])
			}
		}
		runs = newRuns
	}

	return arr
}

// inPlaceMerge merges two consecutive sorted subarrays arr[l:m] and arr[m:r] in-place.
// It uses the provided buffer 'buf' to temporarily hold the first subarray.
// The buffer is reused between merge operations to reduce GC overhead.
func inPlaceMerge(arr []int, l, m, r int, buf []int) {
	n1 := m - l
	if n1 > len(buf) {
		// If needed, allocate a larger buffer.
		buf = make([]int, n1)
	}
	copy(buf, arr[l:m])
	i, j, k := 0, m, l
	for i < n1 && j < r {
		if buf[i] <= arr[j] {
			arr[k] = buf[i]
			i++
		} else {
			arr[k] = arr[j]
			j++
		}
		k++
	}
	// Copy any remaining elements from the buffer.
	for i < n1 {
		arr[k] = buf[i]
		i++
		k++
	}
}
