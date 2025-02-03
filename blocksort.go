package sortcomparison

/*
BlockSort Implementation (Cache-Efficient Hybrid Sort)

Time Complexity:
  - Overall:  O(n log n)
  - Blocks:   O(n * sqrt(n)) - insertion sort on sqrt(n) sized blocks
  - Merging:  O(n log n) - similar to merge sort

Space Complexity:
  - O(n) - requires temporary array for merging
  - O(1) - for block sorting phase (in-place)

Implementation Notes:
  - Block size chosen as sqrt(n) for optimal cache performance
  - Combines insertion sort (small ranges) with merge sort (large ranges)
  - Cache-friendly due to localized memory access patterns
  - Stable sort - maintains relative order of equal elements
  - Good performance on partially sorted data
*/
func BlockSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Determine block size - sqrt(n) is generally a good choice
	blockSize := 1
	for blockSize*blockSize < n {
		blockSize++
	}

	// Sort individual blocks using insertion sort
	for start := 0; start < n; start += blockSize {
		end := start + blockSize
		if end > n {
			end = n
		}
		insertionSortRange(arr, start, end)
	}

	// Merge sorted blocks
	temp := make([]int, n)
	for size := blockSize; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := left + size
			right := left + 2*size
			if mid > n {
				mid = n
			}
			if right > n {
				right = n
			}
			if mid < right {
				mergeBlocks(arr, temp, left, mid, right, func(a, b int) bool { return a < b })
			}
		}
	}
}

func mergeBlocks(arr, temp []int, left, mid, right int, less func(a, b int) bool) {
	// Copy to temporary array
	for i := left; i < right; i++ {
		temp[i] = arr[i]
	}

	i := left
	j := mid
	k := left

	// Merge back to original array
	for i < mid && j < right {
		if less(temp[i], temp[j]) {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}

	// Copy remaining elements
	for i < mid {
		arr[k] = temp[i]
		k++
		i++
	}
	// Note: remaining elements from right side are already in place
}
