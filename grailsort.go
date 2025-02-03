package sortcomparison

import (
	"math"
)

// GrailSort sorts a slice of integers in ascending order using a simplified GrailSort algorithm.
// GrailSort is designed to be efficient, GC friendly, and idiomatic Go code.
// This implementation uses a bottom-up merge sort structure with a fixed-size temporary buffer
// (of size O(√n)) for merging small runs. When a run is larger than the buffer, it falls back to a
// standard merge that allocates a temporary slice. The algorithm is stable and minimizes garbage
// collection overhead by reusing the fixed buffer in most merge operations.
//
// Time Complexity:
//   - Worst-case: O(n log n)
//
// Space Complexity:
//   - O(√n) auxiliary space for the fixed buffer (with rare larger temporary allocations).
func GrailSort(arr []int) {
	n := len(arr)
	if n < 32 {
		insertionSort(arr)
		return
	}

	// Choose fixed buffer size approximately equal to √n.
	blockSize := int(math.Sqrt(float64(n)))
	if blockSize < 1 {
		blockSize = 1
	}
	buf := make([]int, blockSize)

	// Bottom-up merge sort using a fixed buffer for merging.
	for width := 1; width < n; width *= 2 {
		for i := 0; i < n; i += 2 * width {
			mid := i + width
			if mid > n {
				mid = n
			}
			end := i + 2*width
			if end > n {
				end = n
			}
			mergeWithBuffer(arr, i, mid, end, buf)
		}
	}
}

// mergeWithBuffer merges two consecutive sorted subarrays: arr[left:mid] and arr[mid:right].
// If the left subarray fits in the provided buffer, it is used to minimize dynamic allocations.
// Otherwise, the function falls back to allocating a temporary slice for the left run.
func mergeWithBuffer(arr []int, left, mid, right int, buf []int) {
	leftLen := mid - left
	if leftLen <= len(buf) {
		// Use the fixed buffer to hold the left run.
		copy(buf, arr[left:mid])
		i, j, k := 0, mid, left
		for i < leftLen && j < right {
			if buf[i] <= arr[j] {
				arr[k] = buf[i]
				i++
			} else {
				arr[k] = arr[j]
				j++
			}
			k++
		}
		// Copy any remaining elements from buf.
		for i < leftLen {
			arr[k] = buf[i]
			i++
			k++
		}
	} else {
		// Fallback: allocate a temporary slice for the left run.
		temp := make([]int, leftLen)
		copy(temp, arr[left:mid])
		i, j, k := 0, mid, left
		for i < leftLen && j < right {
			if temp[i] <= arr[j] {
				arr[k] = temp[i]
				i++
			} else {
				arr[k] = arr[j]
				j++
			}
			k++
		}
		for i < leftLen {
			arr[k] = temp[i]
			i++
			k++
		}
	}
}
