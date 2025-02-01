package sortcomparison

/*
SimplePancakeSort Implementation

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²)
  - Best:    O(n²)

Space Complexity:
  - O(1) - only requires variables for indexing
  - In-place sorting algorithm

Implementation Notes:
  - Based on flipping prefixes of the array
  - Not stable - does not preserve order of equal elements
  - Each flip reverses a portion of the array
  - Useful for sorting problems with reversal operations
  - Real-world application in parallel processor networks
*/
func SimplePancakeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for size := n; size > 1; size-- {
		// Find index of maximum element in arr[0:size]
		maxIdx := 0
		for i := 1; i < size; i++ {
			if arr[i] > arr[maxIdx] {
				maxIdx = i
			}
		}

		if maxIdx == size-1 {
			continue
		}

		// If max element is not at beginning, flip to move it there
		if maxIdx != 0 {
			flip(arr, maxIdx)
		}

		// Flip again to move max element to its final position
		flip(arr, size-1)
	}
}

func flip(arr []int, k int) {
	for i, j := 0, k; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
